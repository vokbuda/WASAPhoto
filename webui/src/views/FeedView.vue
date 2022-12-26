
<script>
import {adjustNumber} from '../helpers/adjustNumber.js'
import {convertToString} from '../helpers/convertToString.js'
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
			some_data: null,
			feed_posts:[]
		}
	},
	methods: {
		async getLastPosts(){
			this.header=sessionStorage.getItem("token")

		
		
			var userid=sessionStorage.getItem("userid")


			
		
			await this.$axios.get('/posts?offset=0&userid='+userid,{
				headers:{
					'Authorization':this.header,
				
				}




			})
				.then(response => {
					// Handle response
					
					
					if (response.status==200){
						this.feed_posts=response.data
						console.log(this.feed_posts)
						this.feed_posts.forEach((element, index) => {
							element.quantityLikes=adjustNumber(element.quantityLikes)
							element.quantityDislikes=adjustNumber(element.quantityDislikes)
						});
						

					}else{
						this.feed_posts=[]
					}
					
					
				})
				.catch(err => {
					// Handle errors
					console.error(err);
				
				});

		},
		async refresh() {
			this.loading = true;
			this.errormsg = null;
			/*
			try {
				
				let response = await this.$axios.get("/");
				this.some_data = response.data;
			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;*/
		},
	},
	mounted() {
		this.getLastPosts()
	}
}

</script>
<style scoped>
.card-img-top {
    width: 100%;
    height: 15vw;
    object-fit: cover;
}
.in_a_row{
	display:flex;
	justify-content:space-between;
	align-items: center
}
.avatar {
  width: 30px;
  border-radius: 50%;
}
.avatar-bordered {
  box-shadow: 0 1px 2px rgba(0,0,0,0.2);
  border: white 1px solid;
}
.avatar-large {
  width: 50px;
}
#gallery{
	margin-top: 30px
}



</style>


<template>
	<div>
		
		
		<section id="gallery">
		<div class="container">
			<div class="row">
			<div class="col-lg-4 mb-4" v-for="(post, index) in feed_posts" :key="index">
		<div class="card">
			<img v-bind:src="'data:image/jpeg;base64,'+post.image" alt="" class="card-img-top">
			
			<div class="card-body">
				<div class="in_a_row">
				<div>
				<img class="card-user avatar avatar-large" src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fsafeharborpartners.com%2Fwp-content%2Fuploads%2Fshutterstock_169562684-449x375.jpg&f=1&nofb=1&ipt=fe4b42d35bb3eb2cf3d88d1eb7ebcb7e883e15736e51a2db2367cbf4f9eca201&ipo=images">
				
				<h5 class="card-title">{{post.authorname}}</h5>
				</div>
				<btn @click="this.choosePost(post)" data-bs-toggle="modal" data-bs-target="#updatePostModal" style="margin-right:2px"><i style="font-size:2em;" class="bi bi-pencil-fill"></i></btn>
				<btn @click="this.choosePost(post)" data-bs-toggle="modal" data-bs-target="#deletePostModal"><i style="font-size:2em;" class="bi bi-trash-fill"></i></btn>
				</div>
				<p class="card-text">{{post.text}}</p>
				<div class="in_a_row">
				<btn @click="this.sendToComments(post.postid)" class="btn btn-outline-success btn-sm">Comments</btn>

				<btn v-if="post.currentemotion!=1" @click="this.likePost(post)" class="btn btn-outline-danger btn-sm"><i class="bi bi-hand-thumbs-up"></i></btn>
				<btn v-else @click="this.deletePostLike(post)" class="btn btn-outline-danger btn-sm"><i class="bi bi-hand-thumbs-up-fill"></i></btn>
				<h5>{{post.quantityLikes}}</h5>

				<btn v-if="post.currentemotion!=-1" @click="this.dislikePost(post)" class="btn btn-outline-danger btn-sm"><i class="bi bi-hand-thumbs-down"></i></btn>
				<btn v-else @click="this.deletePostDislike(post)" class="btn btn-outline-danger btn-sm"><i class="bi bi-hand-thumbs-down-fill"></i></btn>
				<h5>{{post.quantityDislikes}}</h5>
				</div>
				
			</div>
			</div>
			</div>
		
			
		</div>
		</div>
		</section>

		<!--
		<div style="margin-top: 30px;">
        <div class="card-group">
            <div class="card">
                <img src="..." class="card-img-top" alt="...">
                <div class="card-body">
                <h5 class="card-title">Card title</h5>
                <p class="card-text">This is a wider card with supporting text below as a natural lead-in to additional content. This content is a little bit longer.</p>
                <p class="card-text"><small class="text-muted">Last updated 3 mins ago</small></p>
                </div>
            </div>
            <div class="card">
                <img src="..." class="card-img-top" alt="...">
                <div class="card-body">
                <h5 class="card-title">Card title</h5>
                <p class="card-text">This card has supporting text below as a natural lead-in to additional content.</p>
                <p class="card-text"><small class="text-muted">Last updated 3 mins ago</small></p>
                </div>
            </div>
            <div class="card">
                <img src="..." class="card-img-top" alt="...">
                <div class="card-body">
                <h5 class="card-title">Card title</h5>
                <p class="card-text">This is a wider card with supporting text below as a natural lead-in to additional content. This card has even longer content than the first to show that equal height action.</p>
                <p class="card-text"><small class="text-muted">Last updated 3 mins ago</small></p>
                </div>
            </div>
        </div>
        </div>-->

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
@import "https://cdn.jsdelivr.net/npm/bootstrap-icons@1.10.2/font/bootstrap-icons.css";



</style>
