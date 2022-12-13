<script setup>
import { RouterLink, RouterView } from 'vue-router'

</script>
<script>
export default {
	data: function() {
		return {
			username:"",
			password:""
		}
	},
	
	methods: {
		session(username, password) {
			console.log(username,password)
			const sessionData = JSON.stringify({ "username": username,"password":password,"token":"" });
			this.$axios.post('http://192.168.126.129:3000/session', sessionData
			)
			.then(function (response) {
				console.log(response);
			})
			.catch(function (error) {
				console.log(error);
			});

			/*
			try {
				
				let response = await this.$axios.post("/session")
				// there should be data inside of component

				this.some_data = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}*/
			
			
		},
	},
    mounted() {
		console.log("Entrance inside of register vue check inside current components")
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
			<div class="mb-3">
				<label class="form-label" >Password</label>
				<input type="password" class="form-control" v-model="password">
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
