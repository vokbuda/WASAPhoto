
<script>
import router from '../router'
import FollowerProfileItem from '../components/FollowerProfileItem.vue'
import LoadingSpinner from '../components/LoadingSpinner.vue'
export default {
	data: function() {
		return {
			errormsg: null,
			
			some_data: null,
			searchedUsername: "",
			answer:"",
			header:"",
			usersArray:[],
			offset:0,
			loading:true
		}
	},
	components:{
		FollowerProfileItem,
		LoadingSpinner
	},
	methods: {
		async subscribe(profile){
			this.userid=sessionStorage.getItem("userid")
			const postData = JSON.stringify({ "followingUserid": parseInt(this.userid),"followedUserid":profile.userid });
			
			
			
			await this.$axios.put('/profiles/'+this.userid+'/subscribe/'+profile.userid, postData,{
				headers:{
					"Authorization":sessionStorage.getItem("token"),
					"Content-Type":'application/json'
				}
			}
			)
			.then((response)=> {
				if(response.status==204){
					profile.currentFollow=true
					


				}
				
				
				
				
			})
			.catch(function (error) {
				console.log(error);
			});

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
					
					
					
					
					//document.getElementById("closeModalPostCreate").click()
					//this.createdPost()
					
				}
			})
			.catch(function (error) {
				console.log(error);
			});

		},
		async ban(profile){
			this.userid=sessionStorage.getItem("userid")
			
			const postData = JSON.stringify({ "banningUserid": parseInt(this.userid),"bannedUserid":profile.userid });
			
			
			
			await this.$axios.put('/profiles/'+this.userid+'/banuser/'+profile.userid, postData,{
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
		gotoProfile(userid){
			router.push("/profiles/"+userid)

			
		},
		async getFollowers() {
			this.header=sessionStorage.getItem("token")
            // then u should have data for the following users below 
            // /profiles/:userid/following
			
			try {
				await fetch(__API_URL__+'/profiles/'+this.$route.params.userid+'/followers?offset='+this.offset,
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
						this.loading=false
						
						
					})

				
				// this.answer = (await res.json()).answer
				
				
			} catch (error) {
				this.answer = 'Error! Could not reach the API. ' + error
			}
    	},
		

	},
	watch: {
    // whenever question changes, this function will run
		searchedUsername(newSearchedUsername, oldSearchedUsername) {
			
			this.getAnswer(newSearchedUsername)
			
      
			//this.searchUsers(newSearchedUsername)
			//console.log("data for check inside")

			// here u should have data inside of current compone
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
		
        
                 

		
		<div v-if="!loading">
		<div style="margin-top: 30px;">
        <div class="card-group" v-if="usersArray!=null">
            <div class="container-fluid">
			<div class="row">
				<FollowerProfileItem class="col-md-4 py-2" v-for="(person, index) in usersArray" :key="index"
				:person="person"
				@subscribe="subscribe(person)"
				@unban="unban(person)"
				@ban="ban(person)"
				@unsubscribe="unsubscribe(person)"
				@gotoProfile="gotoProfile(person.userid)"
				>
					
					
				</FollowerProfileItem>
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
		<div>
			
		</div>
        </div>
		
		<div v-if="usersArray.length==0" style="text-align:center"><h2>There is no followers</h2></div>

	
	</div>
	<div v-else>
		<LoadingSpinner>
		</LoadingSpinner>
		

	</div>
        

	<div v-observe-visibility="getFollowers"></div>
	</div>
	
</template>

<style>
</style>
