
<script>
import router from '../router'
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			some_data: null,
			searchedUsername: "",
			answer:"",
			header:"",
			usersArray:[],
			offset:0
		}
	},
	methods: {
		keyupMetho(){

		},
		gotoProfile(userid){
			router.push("/profiles/"+userid)

			
		},
		unsubscribe(profile){
			this.userid=sessionStorage.getItem("userid")
			const postData = JSON.stringify({ "followingUserid": parseInt(this.userid),"followedUserid":profile.userid });
			
			
			
			
			fetch(__API_URL__+'/profiles/'+this.userid+'/subscribe/'+profile.userid,{
				method:'DELETE',
				headers:{
					"Authorization":sessionStorage.getItem("token")
				},
				body:postData
			}
			)
			.then((response)=> {
				
				if (response.status==204){
					profile.currentFollow=false
					this.usersArray = this.usersArray.filter(function(el) { return el.userid !=profile.userid; });

					
					
					
					
					//document.getElementById("closeModalPostCreate").click()
					//this.createdPost()
					
				}
			})
			.catch(function (error) {
				console.log(error);
			});

		},
		
		ban(profile){
			this.userid=sessionStorage.getItem("userid")
			
			const postData = JSON.stringify({ "banningUserid": parseInt(this.userid),"bannedUserid":profile.userid });
			
			
			
			this.$axios.put('/profiles/'+this.userid+'/banuser/'+profile.userid, postData,{
				headers:{
					"Authorization":sessionStorage.getItem("token"),
					"Content-Type":'application/json'
				}
			}
			)
			.then((response)=> {
				if(response.status==204){
					profile.currentBan=true
				}
			})
			.catch(function (error) {
				console.log(error);
			});
			

		},
		unban(profile){
			this.userid=sessionStorage.getItem("userid")
			const postData = JSON.stringify({ "banningUserid": parseInt(this.userid),"bannedUserid":profile.userid });
			fetch(__API_URL__+'/profiles/'+this.$route.params.userid+'/banuser/'+profile.userid,{
				method:'DELETE',
				headers:{
					"Authorization":sessionStorage.getItem("token")
				},
				body:postData
			}
			)
			.then((response)=> {
				
				if (response.status==204){
					profile.currentBan=false
				}
			})
			.catch(function (error) {
				console.log(error);
			});
			

		},
		async getFollowing() {
			this.header=sessionStorage.getItem("token")
            // then u should have data for the following users below 
            // /profiles/:userid/following
			
			try {
				
				await fetch(__API_URL__+'/profiles/'+this.$route.params.userid+'/following?offset='+this.offset,
				{headers:{
					'Authorization':this.header,
				
				}
				}).then((response)=>{
					this.offset+=10
					
					return response.json();
					}).then((data)=>{
					
						if(data){
							
							this.usersArray=this.usersArray.concat(data)
						}else{
							this.offset-=10
						}
						
						
						
					})

				
				// this.answer = (await res.json()).answer
				
				
			} catch (error) {
				this.answer = 'Error! Could not reach the API. ' + error
			}
    	},
		

	},
	
	mounted(){
        
		// here u should have some additional data, check it inside of current component
	}
	
}
</script>
<style scoped>
.profile-card-4 {
    max-width: 300px;
    background-color: #FFF;
    border-radius: 5px;
    box-shadow: 0px 0px 25px rgba(0, 0, 0, 0.1);
    overflow: hidden;
    position: relative;
    margin: 10px auto;
    cursor: pointer;
}

.profile-card-4 img {
    transition: all 0.25s linear;
}

.profile-card-4 .profile-content {
    position: relative;
    padding: 15px;
    background-color: #FFF;
}

.profile-card-4 .profile-name {
    font-weight: bold;
    position: absolute;
    left: 0px;
    right: 0px;
    top: -70px;
    color: #FFF;
    font-size: 17px;
}

.profile-card-4 .profile-name p {
    font-weight: 600;
    font-size: 13px;
    letter-spacing: 1.5px;
}

.profile-card-4 .profile-description {
    color: #777;
    font-size: 12px;
    padding: 10px;
}

.profile-card-4 .profile-overview {
    padding: 15px 0px;
}

.profile-card-4 .profile-overview p {
    font-size: 10px;
    font-weight: 600;
    color: #777;
}

.profile-card-4 .profile-overview h4 {
    color: #273751;
    font-weight: bold;
}

.profile-card-4 .profile-content::before {
    content: "";
    position: absolute;
    height: 20px;
    top: -10px;
    left: 0px;
    right: 0px;
    background-color: #FFF;
    z-index: 0;
    transform: skewY(3deg);
}

.profile-card-4:hover img {
    transform: rotate(5deg) scale(1.1, 1.1);
    filter: brightness(110%);
}
</style>

<template>
	<div>
		
        
                 

		
		<div>
		<div style="margin-top: 30px;">
        <div class="card-group" v-if="usersArray!=null">
            <div class="container-fluid">
			<div class="row">
				<div class="col-md-4 py-2" v-for="(person, index) in usersArray" :key="index">
					
					<div  class="profile-card-4 text-center">
						<div @click="this.gotoProfile(person.userid)" v-if="person.avatar"><img :src="'data:image/jpeg;base64,'+person.avatar" class="img img-responsive"></div>
						<div @click="this.gotoProfile(person.userid)" v-else><img src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fsafeharborpartners.com%2Fwp-content%2Fuploads%2Fshutterstock_169562684-449x375.jpg&f=1&nofb=1&ipt=fe4b42d35bb3eb2cf3d88d1eb7ebcb7e883e15736e51a2db2367cbf4f9eca201&ipo=images" class="img img-responsive"></div>
						<div class="profile-content">
							<div class="profile-name" @click="this.gotoProfile(person.userid)">{{person.username}}</div>
						</div>
                        <div v-if="person.mine" class="row">
								
								<div class="col-xs-4">
									<button @click="this.unsubscribe(person)" style="margin-bottom:10px" type="button" class="btn btn-warning">unfollow</button>
								</div>
                                <div v-if="!person.currentBan" class="col-xs-4">
                                    <button @click="this.ban(person)" style="margin-bottom:10px" type="button" class="btn btn-danger">ban</button>
                                    
									
								</div>
                                <div v-else class="col-xs-4">
                                    <button @click="this.unban(person)" style="margin-bottom:10px" type="button" class="btn btn-danger">unban</button>
                                    
									
								</div>
								
							</div>
					</div>
                    
				</div>
				
				<!--
				<div class="col-md-4 py-2" v-for="(person, index) in usersArray" :key="index">
				<div class="card h-100">
					<img class="card-img-top" :src="person.avatar" alt="card image collar">
					<div class="card-body d-flex flex-column align-items-start">
					<h5 class="card-title">Product {{index}}</h5>
					<p class="card-text">Product {{index}} - {{person.username}}</p>
					<button v-on:click="addProductToCart(product)" class="btn btn-primary mt-auto">Add To Cart</button>
					</div>
				</div>
				</div>-->
			</div>
			</div>
        </div>
		<div v-observe-visibility="this.getFollowing"></div>
		<div>
			
		</div>
        </div>
		<div v-if="!this.usersArray" style="text-align:center"><h2>There is no data</h2></div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
        

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
	
</template>

<style>
</style>
