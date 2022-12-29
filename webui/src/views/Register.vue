<script setup>
import { RouterLink, RouterView } from 'vue-router'
import router from '../router'

</script>
<script>
export default {
	data: function() {
		return {
			username:""
		}
	},
	
	methods: {
		async session(username) {
			
			const sessionData = JSON.stringify({ "username": username});
			var current_token=""
			if(sessionStorage.getItem("token")){
				current_token=sessionStorage.getItem("token")

			}
			
			await this.$axios.post('/session', sessionData,{
				headers:{
					"Authorization":current_token,
					"Content-Type":'application/json'
				}
			}
			)
			.then((response) =>{
				
				if (response.status==200){
					
					sessionStorage.setItem("token",response.data.session)
					sessionStorage.setItem("userid",response.data.uid)
					sessionStorage.setItem("username",username)
					router.push('/welcome')

				}
			})
			.catch(function (error) {
				console.log(error);
			});	
		},
	},
    mounted() {
		
	}
	
}
</script>

<template>
	<div>
	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="/register">WASAPhoto</a>
		<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
	</header>



    <div style="margin-top:50px; margin-left:50px;">
		<form>
			<div class="mb-3">
				<label class="form-label">Username</label>
				<input type="text" class="form-control" aria-describedby="emailHelp" v-model="username">
				<div id="emailHelp" class="form-text" ></div>
			</div>
			
			<button type=button @click="session(username,password)" class="btn btn-success">Submit</button>
		</form>
	</div>

	<div class="container-fluid">
		<div class="row">
			
		</div>
	</div>
	</div>
</template>

<style>
</style>
