<script>
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			some_data: null,
		}
	},
	methods: {
		
	},
	mounted() {
		
		
	}
}
</script>

<template>
	<div>
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">WASAPhoto is an application built for sharing emotions with world</h1>
			
		</div>
		
		<h4>At the moment there had been developed following unique functionalities:</h4>
		<ul class="list-group">
			<li class="list-group-item">Operations with account(change avatar,change username, delete account)</li>
			<li class="list-group-item">Ban and unban users component</li>
			<li class="list-group-item">Operations with posts and comments(create,update, delete)</li>
			<li class="list-group-item">Dislikes on posts and comments</li>
			<li class="list-group-item">Search users</li>
		</ul>


		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
</style>
