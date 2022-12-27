<script setup>
import { RouterLink, RouterView } from 'vue-router'
import router from '../router'

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
		async session(username, password) {
			
			const sessionData = JSON.stringify({ "username": username,"password":password,"token":"" });
			this.$axios.post('http://192.168.126.129:3000/session', sessionData
			)
			.then(function (response) {
				if (response.status==200){
					
					sessionStorage.setItem("token",response.data.session)
					
					router.push('/application')

					
					/*
					localStorage.token=data.token
					router.push({path:'/'})
					*/
					
					// below u set some token inside of session
					


				}
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
		
	}
	
}
</script>

<template>
    <div>
        <center>
            <h1>404 not found</h1>
            <h2>it seems you're in the wrong page</h2>
        </center>
    </div>
</template>

<style>
</style>