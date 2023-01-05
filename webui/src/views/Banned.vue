
<script>
import router from '../router'
import BannedUserItem from '../components/BannedUserItem.vue'
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
		BannedUserItem,
		LoadingSpinner
	},
	methods: {
		subscribe(profile){
			this.userid=sessionStorage.getItem("userid")
			const postData = JSON.stringify({ "followingUserid": parseInt(this.userid),"followedUserid":profile.userid });
			
			
			
			this.$axios.put('/profiles/'+this.userid+'/subscribe/'+profile.userid, postData,{
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
					this.usersArray= this.usersArray.filter(function(el) { return el.userid != profile.userid; });
					
				}
			})
			.catch(function (error) {
				console.log(error);
			});
			

		},
		
		gotoProfile(userid){
			router.push("/profiles/"+userid)

			
		},
		async getBannedUsers() {
			this.header=sessionStorage.getItem("token")
            // then u should have data for the following users below 
            // /profiles/:userid/following
			
			try {
				await fetch(__API_URL__+'/profiles/'+this.$route.params.userid+'/banuser?offset='+this.offset,
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

</style>

<template>
	<div>
		
        
                 

		
		<div v-if="!loading">
		<div style="margin-top: 30px;">
        <div class="card-group" v-if="usersArray!=null">
            <div class="container-fluid">
			<div class="row">
				<BannedUserItem class="col-md-4 py-2" v-for="(person, index) in usersArray" :key="index"
				:person="person"
				@unsubscribe="unsubscribe(person)"
				@subscribe="subscribe(person)"
				@unban="unban(person)"
				@gotoProfile="gotoProfile(person.userid)"
				
				>
				</BannedUserItem>
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
		<div v-if="usersArray.length==0" style="text-align:center"><h2>There is no banned users</h2></div>
		

		
	</div>
	<div v-else>
		<LoadingSpinner>

		</LoadingSpinner>


	</div>

	<div v-observe-visibility="getBannedUsers"></div>
        

	
	</div>
	
</template>

<style>
</style>
