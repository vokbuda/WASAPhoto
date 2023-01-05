<script>
import {adjustNumber} from '../helpers/adjustNumber.js'
export default {
  props: ['comment'],
 
    methods:{
        adjustNumber(data){
            return adjustNumber(data)
        },
        chooseComment(comment){
            this.$emit("choosenComment")
        },
        deleteCommentLike(comment){
            this.$emit("deleteCommentLike")

        },
        deleteCommentDislike(comment){
            this.$emit("deleteCommentDislike")
        },
        commentLike(comment){
            this.$emit("commentLike")
        },
        commentDislike(comment){
            this.$emit("commentDislike")
        },
        goAnotherProfile(){
            this.$emit("goAnotherProfile")
        }

    },
    mounted(){
        console.log("check for current userid below ")
       console.log(this.comment)
        console.log(this.userid==this.comment.authorid)

        
    }    
}
</script>
<style scoped src="@/assets/comment.css">


</style>

<template>
<div>

  <div class="card">
                            <div class="table">
                                
                                <h4 class="card-caption">
	    									<a>{{comment.text}}</a>
	    								</h4>
                                <div class="ftr">
                                    <div class="author">
                                        <a @click="goAnotherProfile()"> <div v-if="!comment.avatar"><img src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fsafeharborpartners.com%2Fwp-content%2Fuploads%2Fshutterstock_169562684-449x375.jpg&f=1&nofb=1&ipt=fe4b42d35bb3eb2cf3d88d1eb7ebcb7e883e15736e51a2db2367cbf4f9eca201&ipo=images" alt="" class="avatar img-raised"> </div>
                                        <div v-else><img v-bind:src="'data:image/jpeg;base64,'+comment.avatar" alt="" class="avatar img-raised"></div>
                                        <span>{{comment.username}}</span>
                                            <div class="ripple-cont">
                                                <div class="ripple ripple-on ripple-out" style="left: 574px; top: 364px; background-color: rgb(60, 72, 88); transform: scale(11.875);"></div>
                                            </div>
                                        </a>
                                    </div>
                                    <div class="stats">
                                    <div v-if="comment.me">     
                                    <btn @click="chooseComment(comment)" data-bs-toggle="modal" data-bs-target="#updateCommentModal"><i style="font-size:1.5em;" class="bi bi-pencil-fill"></i></btn>
                                    <btn @click="chooseComment(comment)" style="margin-left:5px" data-bs-toggle="modal" data-bs-target="#deleteCommentModal"><i style="font-size:1.5em;" class="bi bi-trash-fill"></i></btn>
                                    </div>
                                    <div v-if="comment.currentemotion!=1"><btn @click="commentLike(comment)" class="btn btn-outline-danger btn-sm"><i class="bi bi-hand-thumbs-up"></i></btn></div>
                                    <div v-else><btn @click="deleteCommentLike(comment)" class="btn btn-outline-danger btn-sm"><i class="bi bi-hand-thumbs-up-fill"></i></btn></div>
				{{adjustNumber(comment.quantityLikes)}}

				<div v-if="comment.currentemotion!=-1"><btn @click="commentDislike(comment)" class="btn btn-outline-danger btn-sm"><i class="bi bi-hand-thumbs-down"></i></btn></div>
                <div v-else><btn @click="deleteCommentDislike(comment)" class="btn btn-outline-danger btn-sm"><i class="bi bi-hand-thumbs-down-fill"></i></btn></div>
				{{adjustNumber(comment.quantityDislikes)}}</div>
                                    
                                </div>
                            </div>
                        </div>



        


        


</div>
</template>