
<script>
import router from '../router'
import SearchProfileItem from '../components/SearchProfileItem.vue'
import LoadingSpinner from '../components/LoadingSpinner.vue'




export default {
	components:{
		SearchProfileItem,
		LoadingSpinner

	},
	
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
	methods: {
		async visibilityChanged(){
			
			if(this.searchedUsername){
				await this.loadMorePosts()

			}else{
				this.loading=false
			}

			
			
		},
		async loadMorePosts(){
			this.header=sessionStorage.getItem("token")
			
			this.offset+=10
			
			try {
				await fetch(__API_URL__+'/profiles?username='+this.searchedUsername+'&&offset='+this.offset,
				{headers:{
					'Authorization':this.header,
				
				}
				}).then((response)=>{
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
		
		gotoProfile(userid){
			router.push("/profiles/"+userid)


			
		},
		async getAnswer(username) {
			if(!username){
				this.loading=false
				return
			}
			console.log("Check for some answer inside {%s}",username)
			this.header=sessionStorage.getItem("token")
			this.searchedUsername=username
			
			try {
				await fetch(__API_URL__+'/profiles?username='+username+'&&offset=0',
				{headers:{
					'Authorization':this.header,
				
				}
				}).then((response)=>{
					return response.json();
					}).then((data)=>{
						this.usersArray=data
						this.offset=0
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
			
			if (!newSearchedUsername){
				console.log("are we removing data from here")
				this.usersArray=[]
				this.loading=false
				return
			}
			console.log("check for string newSearched username inside{%s}",newSearchedUsername)
			this.getAnswer(newSearchedUsername)
			
			
			
      
			//this.searchUsers(newSearchedUsername)
			//console.log("data for check inside")

			// here u should have data inside of current compone
    },
  	},
	mounted(){
		this.offset+=10
		
	}
	
}
</script>
<style scoped>

</style>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">Search</h1>
			<!--
			<div class="btn-toolbar mb-2 mb-md-0">
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="refresh">
						Refresh
					</button>
					<button type="button" class="btn btn-sm btn-outline-secondary" @click="exportList">
						Export
					</button>
				</div>
				<div class="btn-group me-2">
					<button type="button" class="btn btn-sm btn-outline-primary" @click="newItem">
						New
					</button>
				</div>
                
			</div>-->
		</div>
        
                 

		<div class="input-group">
			<input class="form-control rounded" v-model="searchedUsername" placeholder="INSERT USERNAME" aria-label="Search" aria-describedby="search-addon" />
			
		</div>
		<div v-if="!loading">
		<div style="margin-top: 30px;">
        <div class="card-group" v-if="usersArray!=null">
            <div class="container-fluid">
			<div class="row" ref="scrollComponent">
				<SearchProfileItem class="col-md-4 py-2" v-for="(person, index) in usersArray" :key="index" :person="person"
				@gotoProfile="gotoProfile(person.userid)"
				>
					
					
				</SearchProfileItem>
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
		<div v-if="usersArray.length==0" style="text-align:center"><h2>There is no data</h2></div>

		

	</div>
	<div v-else>
		<LoadingSpinner></LoadingSpinner>
	</div>
	<div v-observe-visibility="visibilityChanged"></div>
        

		
	</div>
	
</template>

<style>
</style>
