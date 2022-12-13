<script setup>
import { RouterLink, RouterView } from 'vue-router'
</script>
<script>
export default {
    mounted() {
		console.log("Entrance inside of register vue check inside current components")
	}
}
</script>

<template>
	<div>
	<header class="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
		<a class="navbar-brand col-md-3 col-lg-2 me-0 px-3 fs-6" href="/register">RegisterView</a>
		<button class="navbar-toggler position-absolute d-md-none collapsed" type="button" data-bs-toggle="collapse" data-bs-target="#sidebarMenu" aria-controls="sidebarMenu" aria-expanded="false" aria-label="Toggle navigation">
			<span class="navbar-toggler-icon"></span>
		</button>
	</header>



    <div> egisterview component</div>

	<div class="container-fluid">
		<div class="row">
			
		</div>
	</div>
	</div>
</template>

<style>
</style>
