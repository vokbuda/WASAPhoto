
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
			usersArray:[]
		}
	},
	methods: {
		keyupMetho(){

		},
		gotoProfile(userid){
			router.push("/profiles/"+userid)

			
		},
		async getAnswer(username) {
			this.header="Bearer "+sessionStorage.getItem("token")
			
			try {
				await fetch(__API_URL__+'/profiles?username='+username+'&&offset=0',
				{headers:{
					'Authorization':this.header,
				
				}
				}).then((response)=>{
					return response.json();
					}).then((data)=>{
						this.usersArray=data
						console.log(this.usersArray)
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
							<div class="row">
								
								<div class="col-xs-4">
									<div class="profile-overview">
										<p>FOLLOWERS</p>
										<h4>{{person.quantitySubscribers}}</h4></div>
								</div>
								<div class="col-xs-4">
									<div class="profile-overview">
										<p>FOLLOWING</p>
										<h4 >{{person.quantitySubscriptions}}</h4></div>
								</div>
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
