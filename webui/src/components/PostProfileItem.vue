<script>
import {adjustNumber} from '../helpers/adjustNumber.js'
export default {
  props: ['post','profile'],
 
    methods:{
        adjustNumber(data){
            return adjustNumber(data)
        },
        sendToComments(){
            this.$emit("sendToComments")
        },
        choosePost(){
            this.$emit("choosePost")

        },

        deletePostLike(){
            this.$emit("deletePostLike")

        },
        deletePostDislike(){
            this.$emit("deletePostDislike")
        },
        
        likePost(){
            this.$emit("likePost")
        },
        dislikePost(){
            this.$emit("dislikePost")
        }


    },
    mounted(){
        console.log("Below u can check some data for current component")
        console.log(this.post)
        
    }    
}
</script>
<style scoped src="@/assets/postProfile.css">


</style>

<template>
<div>

  <div class="card">
			<img v-bind:src="'data:image/jpeg;base64,'+post.image" alt="" class="card-img-top">
			
			<div class="card-body">
				<div class="in_a_row">
				<div>
				<div v-if="!profile.avatar">
				<img class="card-user avatar avatar-large" src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fsafeharborpartners.com%2Fwp-content%2Fuploads%2Fshutterstock_169562684-449x375.jpg&f=1&nofb=1&ipt=fe4b42d35bb3eb2cf3d88d1eb7ebcb7e883e15736e51a2db2367cbf4f9eca201&ipo=images">
				</div>
				<div v-else>
					<img class="card-user avatar avatar-large" v-bind:src="'data:image/jpeg;base64,'+profile.avatar">

				</div>
				<h5 class="card-title">{{profile.username}}</h5>
				</div>
				<div v-if="profile.me">
				<btn @click="choosePost(post)" data-bs-toggle="modal" data-bs-target="#updatePostModal" style="margin-right:2px"><i style="font-size:2em;" class="bi bi-pencil-fill"></i></btn>
				<btn @click="choosePost(post)" data-bs-toggle="modal" data-bs-target="#deletePostModal"><i style="font-size:2em;" class="bi bi-trash-fill"></i></btn>
				</div>
				</div>
                
				<h3 class="card-text">{{post.text}}</h3>
                <div style="margin-top:40px;">
                <div class="last_data">
				<div class="in_a_row">
                
				<btn @click="sendToComments(post.postid)" class="btn btn-outline-success btn-sm">Comments</btn>
                <div style="margin-left:10px;">
				<btn v-if="post.currentemotion!=1" @click="likePost(post)" class="btn btn-outline-danger btn-sm"><i class="bi bi-hand-thumbs-up"></i></btn>
				<btn v-else @click="deletePostLike(post)" class="btn btn-outline-danger btn-sm"><i class="bi bi-hand-thumbs-up-fill"></i></btn>
                </div>
                <div style="margin-left:10px;">
                <h5>{{post.quantityLikes}}</h5>
                </div>
                <div style="margin-left:10px;">
				<btn v-if="post.currentemotion!=-1" @click="dislikePost(post)" class="btn btn-outline-danger btn-sm"><i class="bi bi-hand-thumbs-down"></i></btn>
				<btn v-else @click="deletePostDislike(post)" class="btn btn-outline-danger btn-sm"><i class="bi bi-hand-thumbs-down-fill"></i></btn>
                </div>
                <div style="margin-left:10px;">
                <h5>{{post.quantityDislikes}}</h5>
                </div>
				</div>
                </div>
                </div>
                
                
                
			</div>
			</div>



        


        


</div>
</template>