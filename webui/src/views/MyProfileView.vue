<script>
import router from '../router/index.js'
import Post from '../entities/Post'
import PostProfileItem from '../components/PostProfileItem.vue'
import {adjustNumber} from '../helpers/adjustNumber.js'
import {convertToString} from '../helpers/convertToString.js'
import LoadingSpinner from '../components/LoadingSpinner.vue'
export default {
	data: function() {
		return {
			errormsg: null,
			loadingPosts:true,
			loadingProfile:true,
			some_data: null,
			me:null,
			switcherSubscription: true,
			switcherBan:true,
			subscriptionText:"follow",
			banText:"ban",
			userid:"",
			profile:{},
			postImage:"",
			postText:"",
			postCreation:true,
			postsProfile:[],
			Password:"",
			newPassword:"",
			newAvatar:"",
			choosenPost:{},
			loaded:false,
			offset:0,
			newUsername:"",
			notificationText:"",
			profileLoaded:false,
			modifyPost:false,
			usernameExists:false
		}
	},
	components:{
		PostProfileItem,
		LoadingSpinner

	},
	
	
	
	methods: {
		goToFollowing(){
			var currentuserid=this.$route.params.userid
			router.push('/profiles/'+currentuserid+'/following')

		},
		goToFollowers(){
			var currentuserid=this.$route.params.userid
			router.push('/profiles/'+currentuserid+'/followers')


		},
		goToBanned(){
			var currentuserid=this.$route.params.userid
			router.push('/profiles/'+currentuserid+'/banuser')
		},
		async ban(){
			this.userid=sessionStorage.getItem("userid")
			const postData = JSON.stringify({ "banningUserid": parseInt(this.userid),"bannedUserid":parseInt(this.$route.params.userid) });
			
			
			
			await this.$axios.put('/profiles/'+this.$route.params.userid+'/banuser/'+this.userid, postData,{
				headers:{
					"Authorization":sessionStorage.getItem("token"),
					"Content-Type":'application/json'
				}
			}
			)
			.then((response)=> {
				if(response.status==204){
					this.profile.currentBan=true
				}
			})
			.catch(function (error) {
				console.log(error);
			});
			

		},
		unban(){
			this.userid=sessionStorage.getItem("userid")
			const postData = JSON.stringify({ "banningUserid": parseInt(this.userid),"bannedUserid":parseInt(this.$route.params.userid) });
			fetch(__API_URL__+'/profiles/'+this.$route.params.userid+'/banuser/'+this.userid,{
				method:'DELETE',
				headers:{
					"Authorization":sessionStorage.getItem("token")
				},
				body:postData
			}
			)
			.then((response)=> {
				
				if (response.status==204){
					this.profile.currentBan=false
				}
			})
			.catch(function (error) {
				console.log(error);
			});
			

		},
		async subscribe(){
			this.userid=sessionStorage.getItem("userid")
			const postData = JSON.stringify({ "followingUserid": parseInt(this.userid),"followedUserid":parseInt(this.$route.params.userid) });
			
			
			
			await this.$axios.put('/profiles/'+this.$route.params.userid+'/subscribe/'+this.userid, postData,{
				headers:{
					"Authorization":sessionStorage.getItem("token"),
					"Content-Type":'application/json'
				}
			}
			)
			.then((response)=> {
				if(response.status==204){
					this.profile.currentFollow=true
					
					this.profile.quantitySubscribers=convertToString(this.profile.quantitySubscribers,1)
					



				}
				
				
				
				
			})
			.catch(function (error) {
				console.log(error);
			});

		},
		unsubscribe(){
			this.userid=sessionStorage.getItem("userid")
			const postData = JSON.stringify({ "followingUserid": parseInt(this.userid),"followedUserid":parseInt(this.$route.params.userid) });
			
			
			
			
			fetch(__API_URL__+'/profiles/'+this.$route.params.userid+'/subscribe/'+this.userid,{
				method:'DELETE',
				headers:{
					"Authorization":sessionStorage.getItem("token")
				},
				body:postData
			}
			)
			.then((response)=> {
				
				if (response.status==204){
					this.profile.currentFollow=false
					this.profile.quantitySubscribers=convertToString(this.profile.quantitySubscribers,-1)
					
					
					
					//document.getElementById("closeModalPostCreate").click()
					//this.createdPost()
					
				}
			})
			.catch(function (error) {
				console.log(error);
			});

		},
		convertToString(data,changed){
			return convertToString(data,changed)
		},
		adjustNumber(data){	
			return adjustNumber(data)
		},
		
		
		async refresh() {
			this.loading = true;
			this.errormsg = null;

			
		},
		handleImageAvatar(e){
			const selected_image=e.target.files[0]
			this.createBase64ImageAvatar(selected_image)

		},

		handleImage(e){
			
			const selected_image=e.target.files[0]
			this.createBase64Image(selected_image)


		},
		async likePost(post){
			this.userid=sessionStorage.getItem("userid")
			const postData = JSON.stringify({ "idPostEmotion": post.postid,"idUser":parseInt(this.userid) });
			
			
			
			await this.$axios.put('/posts/'+post.postid+'/like/'+this.userid, postData,{
				headers:{
					"Authorization":sessionStorage.getItem("token"),
					"Content-Type":'application/json'
				}
			}
			)
			.then((response)=> {
				
				if (response.status==204){
					
					//document.getElementById("closeModalPostCreate").click()
					//this.createdPost()
					
					if(post.currentemotion==-1){
						post.quantityDislikes=convertToString(post.quantityDislikes,-1)
						
					}
					post.quantityLikes=convertToString(post.quantityLikes,1)
					

					post.currentemotion=1
				}
			})
			.catch(function (error) {
				console.log(error);
			});

		},
		deletePostLike(post){
			this.userid=sessionStorage.getItem("userid")
			const postData = JSON.stringify({ "idPostEmotion": post.postid,"idUser":parseInt(this.userid) });
			
			
			
			
			fetch(__API_URL__+'/posts/'+post.postid+'/like/'+this.userid,{
				method:'DELETE',
				headers:{
					"Authorization":sessionStorage.getItem("token")
				},
				body:postData
			}
			)
			.then((response)=> {
				
				if (response.status==204){
					post.currentemotion=0
					
					post.quantityLikes=convertToString(post.quantityLikes,-1)
					
					//document.getElementById("closeModalPostCreate").click()
					//this.createdPost()
					
				}
			})
			.catch(function (error) {
				console.log(error);
			});

		},
		async dislikePost(post){
			this.userid=sessionStorage.getItem("userid")
			const postData = JSON.stringify({ "idPostEmotion": post.postid,"idUser":parseInt(this.userid) });
			
			
			
			await this.$axios.put('/posts/'+post.postid+'/dislike/'+this.userid, postData,{
				headers:{
					"Authorization":sessionStorage.getItem("token"),
					"Content-Type":'application/json'
				}
			}
			)
			.then((response)=> {
				if (response.status==204){
					if(post.currentemotion==1){
						post.quantityLikes=convertToString(post.quantityLikes,-1)
						
					}
					post.quantityDislikes=convertToString(post.quantityDislikes,1)
					post.currentemotion=-1

					
					
					//document.getElementById("closeModalPostCreate").click()
					//this.createdPost()
					
				}
			})
			.catch(function (error) {
				console.log(error);
			});


		},
		deletePostDislike(post){
			this.userid=sessionStorage.getItem("userid")
			const postData = JSON.stringify({ "idPostEmotion": post.postid,"idUser":parseInt(this.userid) });
			
			
			
			fetch(__API_URL__+'/posts/'+post.postid+'/dislike/'+this.userid,{
				method:'DELETE',
				headers:{
					"Authorization":sessionStorage.getItem("token")
				},
				body:postData
			}
			)
			.then((response)=> {
				
				if (response.status==204){
					post.currentemotion=0
					post.quantityDislikes=convertToString(post.quantityDislikes,-1)
					//document.getElementById("closeModalPostCreate").click()
					//this.createdPost()
					
				}
			})
			.catch(function (error) {
				console.log(error);
			});

		},
		
		notification(data,id_element){
			document.getElementById(id_element).click()
			this.notificationText=data
			document.getElementById("liveToast").style.display="block";
			//setTimeout(1500)
			const delay = ms => new Promise(res => setTimeout(res, ms));
			setTimeout(() => {
                document.getElementById("liveToast").style.display="none";
				this.notificationText=""
				this.postText=""
				this.postImage=""
            }, 1500)

			



		},
		async createPost(){
			

			
			const postData = JSON.stringify({ "text": this.postText,"image":this.postImage });
			
			
			await this.$axios.post('/profiles/'+this.$route.params.userid+'/posts', postData,{
				headers:{
					"Authorization":sessionStorage.getItem("token"),
					"Content-Type":'application/json'
				}
			}
			)
			.then((response)=> {
				if (response.status==200){
				
					const post=new Post(this.$route.params.userid,response.data.postid,this.postText,
					'0','0',true,this.postImage,'',0)
					this.modifyPost=true
					
					
					this.postsProfile.unshift(post)
					
					
					this.notification("Post had been created","closeModalPostCreate")
					document.getElementById("inputImage").value = "";
					
				}
			})
			.catch(function (error) {
				console.log(error);
			});

		},
		async updatePost(){
			
			const postData = JSON.stringify({ "text": this.postText,"image":this.postImage });

			
			
			await this.$axios.put('/profiles/'+this.$route.params.userid+'/posts/'+this.choosenPost.postid, postData,{
				headers:{
					"Authorization":sessionStorage.getItem("token"),
					"Content-Type":'application/json'
				}
			}
			)
			.then((response)=> {
				if (response.status==200){
					this.choosenPost.image=this.postImage
					this.choosenPost.text=this.postText
										
					this.modifyPost=true
					this.notification("Post had been updated","closeModalPostUpdate")
				}
			})
			.catch(function (error) {
				console.log(error);
			});

		},
		createBase64Image(fileObject){
			const reader=new FileReader()
			reader.onloadend=()=>{
				this.postImage=reader.result
				
				this.postImage = this.postImage.replace(/^data:image\/[a-z]+;base64,/, "");
				
				
				
			}
			reader.readAsDataURL(fileObject)
			

		

		},
		createBase64ImageAvatar(fileObject){
			const reader=new FileReader()
			reader.onloadend=()=>{
				this.newAvatar=reader.result
				
				this.newAvatar = this.newAvatar.replace(/^data:image\/[a-z]+;base64,/, "");
				
				
				
			}
			reader.readAsDataURL(fileObject)

		},
	
		async getMyPosts(offset){
			
			this.header=sessionStorage.getItem("token")
		
		
			//this.userid=sessionStorage.getItem("userid")


			
		
			await this.$axios.get('/profiles/'+this.$route.params.userid+'/posts?offset='+offset,{
				headers:{
					'Authorization':this.header,
					'Content-Type':'application/json'
				
				}




			})
				.then(response => {
					this.offset+=10
					
					// Handle response
					

					if (response.data!=null && !this.modifyPost){
						
						console.log(response.data)
						
						var current_data=response.data
						current_data.forEach((element, index) => {
							element.quantityLikes=adjustNumber(element.quantityLikes)
							element.quantityDislikes=adjustNumber(element.quantityDislikes)
						});
						this.postsProfile=this.postsProfile.concat(current_data)
						console.log("below u should check some data and then implement the rest!!!!!!")
						console.log(this.postsProfile)
						

					}else{
						this.modifyPost=true
						this.offset-=10
						
					}
					this.loadingPosts=false
					
					
				})
				.catch(err => {
					// Handle errors
					console.error(err);
				
				});
					


			// then here u should setup your posts of your profile

		},
		sendToComments(post){
			router.push('/posts/'+post.postid+'/comments')
			

		},
		async deleteAccount(){
			this.userid=sessionStorage.getItem("userid")
			
			
		
			
			
			fetch(__API_URL__+'/profiles/'+this.userid+'/deleteAccount',{
				method:'DELETE',
				headers:{
					"Authorization":sessionStorage.getItem("token")
				},
				
			}
			)
			.then((response)=> {
				
				if (response.status==204){
					this.notification("Your account had been removed","closeModalDeleteAccount")
					sessionStorage.clear()
					router.push('/dologin')
					
				}
			})
			.catch(function (error) {
				console.log(error);
			});

		},
		async updateUsername(){
			if(!this.newUsername){
				return
			}
			if(this.newUsername.length>16){
				return
			}

			this.header=sessionStorage.getItem("token")
			const postData = JSON.stringify({ "newValue": this.newUsername});
			
			
			await this.$axios.patch('/profiles/'+this.$route.params.userid+'/changeUsername',postData,{
				headers:{
					'Authorization':this.header,
					'Content-Type':'application/json'
				
				}




			})
				.then(response => {
					if(response.status=='200'){
						this.notification("Username had been updated","closeModalUsernameUpdate")
						this.profile.username=this.newUsername
						sessionStorage.setItem("username",this.newUsername)


					}
				})
				.catch(err => {
					// Handle errors
					
					this.usernameExists=true
					setTimeout(() => {
						this.usernameExists=false
					}, "1500")
					
				
				});

		},
		async updateAvatar(){
			this.header=sessionStorage.getItem("token")
			const postData = JSON.stringify({ "newValue": this.newAvatar });
			
			
			await this.$axios.patch('/profiles/'+this.$route.params.userid+'/changeAvatar',postData,{
				headers:{
					'Authorization':this.header,
					'Content-Type':'application/json'
				
				}




			})
				.then(response => {
					if(response.status=='200'){
						this.notification("Avatar had been updated","closeModalAvatarUpdate")
						this.profile.avatar=this.newAvatar
						sessionStorage.setItem("avatar",this.newAvatar)

					}
				})
				.catch(err => {
					// Handle errors
					console.error(err);
				
				});

		},
		
		async deletePost(){
			this.userid=sessionStorage.getItem("userid")
			
			
			
			var choosenId=this.choosenPost.postid
			
			fetch(__API_URL__+'/profiles/'+this.userid+'/posts/'+choosenId,{
				method:'DELETE',
				headers:{
					"Authorization":sessionStorage.getItem("token")
				}
			}
			)
			.then((response)=> {
				
				if (response.status==204){
					// closeModalDeleteAccount
					this.postsProfile = this.postsProfile.filter(function(el) { return el.postid != choosenId; });
					this.notification("Post had been removed","closeModalDeletePost")
					this.modifyPost=true
					
					
					
				}
			})
			.catch(function (error) {
				console.log(error);
			});
			

		},
		choosePost(post){
			this.choosenPost=post
			this.postText=this.choosenPost.text
			this.postImage=this.choosenPost.image
			
			
			

		},
		async visibilityChanged(){
			await this.getMyPosts(this.offset)
			
			
			
		},
		
		async getMyProfile(){
		
		
			this.header=sessionStorage.getItem("token")
		
		
			


			
		
			await this.$axios.get('/profiles/'+this.$route.params.userid,{
				headers:{
					'Authorization':this.header,
					'Content-Type':'application/json'
				
				}




			})
				.then(response => {
					// Handle response
				
					this.profile=response.data
					
					this.loadingProfile=false
					
					this.profile.quantitySubscriptions=this.adjustNumber(this.profile.quantitySubscriptions)
					this.profile.quantitySubscribers=this.adjustNumber(this.profile.quantitySubscribers)
					
					
				})
				.catch(err => {
					// Handle errors
					console.error(err);
				
				});
					},
		
	},
	/*
	watch: {
	// when redirect to new category_name, this will be callback
		'$route.params.userid': {
			handler: function(userid) {
				router.push('/profiles/'+userid)
			},
			
		}
	},*/
	
	
	created() {
		this.profileLoaded=false
		this.postCreation=true
		this.getMyProfile()
		
		
		
	}
}
</script>

<style scoped>
body{
  min-height:100vh;
  min-width:100vw;
  
  display:flex;
  justify-content:center;
  align-items:center;
}

#header{
  width:50vw;
  max-width:400px;
  min-width:300px;
  height:70px;
  
  transition: all .08s linear;
}

#profile{
  width:50vw;
  max-width:400px;
  min-width:300px;
  height:160px;
  background:white;
  position:relative;
  box-sizing:border-box;
  padding-top:40px;
  padding-left:25px;
  font-family: 'Open Sans', sans-serif;
}
.image{
    
    
    
    height:80px;
    
    
}
.name{
    font-size:1.3rem;
    font-weight:500;
    color:#444;
}
.nickname{
    color:#888;
    font-size:0.9rem;
    padding-bottom:7px;
  }
.location{
    color:#555;
    font-size:0.9rem;
    padding-left:0px;
    position:relative;
    left:-5px;
}
.material-icons{
      position:relative;
      top:3px;
      font-size:1rem;
}
.absolution{
	position: absolute;
	top:-55px;
    left: 20px;
}

img{
      width:inherit;
      height:inherit;
      border-radius:8px;
    }

.shadow{
  box-shadow: 0px 0px 10px 1px rgba(black,0.5);
}

.overflow{
  overflow:hidden;
}

.following,
.followers,
.banned{
  font-family: 'Open Sans', sans-serif;
  font-size:0.9rem;
  /*color:#bbb;*/
  color:orange;
  font-weight:300;
}
.banned{
  float:right;
  padding-right:30px;

}
.followers{
	margin-left: 10px;
   padding-right:30px;
}
.banned,.followers,.following:hover{
	cursor:pointer
}
.bottom{
  margin-top:10px;
}

.count{
  color:black;
  font-family: 'Open Sans', sans-serif;
  font-size:0.9rem;
  font-weight:700;
}
.in_a_row{
	display:flex;
	justify-content:space-between;
	
	
}
.row{
	display:flex;
	align-items: stretch;
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
a:hover{
	cursor:pointer
}

#element1 {display:inline-block;margin-right:10px;} 
#element2 {display:inline-block;} 

.card-img-top {
    width: 100%;
    height: 15vw;
    object-fit: cover;
}
.img-profile{
	width:30px
}

.fab-container {
  
  position: fixed;
  
	z-index: 1;
  bottom: 30px; right: 30px;
  
}

</style>

<template>
<div>
	<div>

<div id="looker" class="position-fixed bottom-0 end-0 p-3" style="z-index: 11">
  <div id="liveToast" class="toast hide" role="alert" aria-live="assertive" aria-atomic="true">
    <div class="toast-header">
      
      <strong class="me-auto">WASAPhoto</strong>
      
      
    </div>
    <div class="toast-body">
      {{notificationText}}
    </div>
  </div>
</div>
		<div v-if="!loadingProfile" class="shadow overflow">
		<div id="header"></div>
		<div id="profile">
			<div class="absolution">
			<div id="element1">
			<div v-if="profile.me">
			<div data-bs-toggle="modal" data-bs-target="#imageModal" v-if="profile.avatar">
				<img class="image" v-bind:src="'data:image/jpeg;base64,'+profile.avatar">
			</div>
			


			<div data-bs-toggle="modal" data-bs-target="#imageModal" v-else><img class="image" src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fsafeharborpartners.com%2Fwp-content%2Fuploads%2Fshutterstock_169562684-449x375.jpg&f=1&nofb=1&ipt=fe4b42d35bb3eb2cf3d88d1eb7ebcb7e883e15736e51a2db2367cbf4f9eca201&ipo=images"></div>
			<!--<img src="https://a4-images.myspacecdn.com/images03/2/85a286a4bbe84b56a6d57b1e5bd03ef4/300x300.jpg" alt="" />-->
			</div>
			<div v-else>
				<div  v-if="profile.avatar">
				<img class="image" v-bind:src="'data:image/jpeg;base64,'+profile.avatar">
			</div>
			


			<div  v-else><img class="image" src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fsafeharborpartners.com%2Fwp-content%2Fuploads%2Fshutterstock_169562684-449x375.jpg&f=1&nofb=1&ipt=fe4b42d35bb3eb2cf3d88d1eb7ebcb7e883e15736e51a2db2367cbf4f9eca201&ipo=images"></div>
			
				

			</div>
			</div>
			<div v-if="profile.me" id="element2">
				<btn data-bs-toggle="modal" data-bs-target="#deleteAccountModal"><i style="font-size:2em;" class="bi bi-trash3-fill"></i></btn>
			</div>
			
			</div>
			<div class="name">
			{{profile.username}}
			<btn v-if="profile.me" data-bs-toggle="modal" data-bs-target="#usernameModal"><i class="bi bi-pencil-square"></i></btn>
			
			
			<div class="modal fade" id="usernameModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
			<div class="modal-dialog">
				<div class="modal-content">
				<div class="modal-header">
					<h5 class="modal-title" id="exampleModalLabel">Setup new username</h5>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				<div class="modal-body">
					<form>
					
					<div class="mb-3">
						<label for="message-text" class="col-form-label">New username:</label>
						<input v-model="newUsername" type="text" class="form-control" id="message-text">
					</div>
					<div v-if="usernameExists"><p class="text-center text-danger">Username already exists</p></div>
					</form>
				</div>
				<div class="modal-footer">
					<button id="closeModalUsernameUpdate" type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
					<div v-if="newUsername">
						<div v-if="newUsername.length<16 && newUsername.length>0">
					<button @click="updateUsername()" type="button" class="btn btn-warning">Save updates</button>
						</div>
					</div>
				</div>
				</div>
			</div>
			</div>
			<div v-if="postCreation" class="modal fade" id="postModal" tabindex="-1" aria-labelledby="exampleModalLabel" >
			<div class="modal-dialog">
				<div class="modal-content">
				<div class="modal-header">
					<h5 class="modal-title" id="exampleModalLabel">Write post below</h5>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				<div class="modal-body">
					<form>
					
					<div class="mb-3">
						<div class="input-group mb-3">
						<input type="file" @change="handleImage" accept="image/*" class="form-control" id="inputImage">
						
						</div>
					</div>
					<div>
						<img v-if="postImage" class="image" v-bind:src="'data:image/jpeg;base64,'+this.postImage">
						<img v-else class="image" src="../images/nophoto.jpg">

					</div>
					<div class="mb-3">
						<label for="message-text" class="col-form-label">Text:</label>
						<input v-model="postText" type="text" class="form-control" >
					</div>
					</form>
				</div>
				<div v-if="!postText||!postImage"><p class="text-center text-danger">u should insert data</p></div>
				<div v-else>
				<div class="modal-footer">
					<button id="closeModalPostCreate" type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
					<button type="button" @click="createPost()" class="btn btn-success">Create</button>
					
				</div>
				</div>
				</div>
			</div>
			</div>
			<div v-if="postCreation" class="modal fade" id="updatePostModal" tabindex="-1" aria-labelledby="exampleModalLabel" >
			<div class="modal-dialog">
				<div class="modal-content">
				<div class="modal-header">
					<h5 class="modal-title" id="exampleModalLabel">Update Post</h5>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				<div class="modal-body">
					<form>
					
					<div class="mb-3">
						<div class="input-group mb-3">
						<input type="file" @change="handleImage" accept="image/*" class="form-control" id="inputGroupFile02">
						
						</div>
					</div>
					<div>
						<img class="image" v-bind:src="'data:image/jpeg;base64,'+this.postImage">
					</div>
					<div class="mb-3">
						<label for="message-text" class="col-form-label">Text:</label>
						<input v-model="postText" type="text" class="form-control" id="message-text">
					</div>
					</form>
				</div>
				<div v-if="!postText||!postImage"><p class="text-center text-danger">u should insert data</p></div>
				<div v-else>
				<div class="modal-footer">
					<button id="closeModalPostUpdate" type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
					<button type="button" @click="updatePost()" class="btn btn-warning">Update</button>
				</div>
				</div>
				</div>
			</div>
			</div>
			
			<div class="modal fade" id="deleteAccountModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
			<div class="modal-dialog">
				<div class="modal-content">
				<div class="modal-header">
					<h5 class="modal-title" id="exampleModalLabel">Delete Account</h5>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				
				<div class="modal-footer">
					<button id="closeModalDeleteAccount" type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
					<button @click="deleteAccount()" type="button" class="btn btn-danger">Confirm</button>
				</div>
				</div>
			</div>
			</div>
			<div class="modal fade" id="deletePostModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
			<div class="modal-dialog">
				<div class="modal-content">
				<div class="modal-header">
					<h5 class="modal-title" id="exampleModalLabel">Delete Post?</h5>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				
				<div class="modal-footer">
					<button id="closeModalDeletePost" type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
					<button @click="deletePost()" type="button" class="btn btn-danger">Delete Post</button>
				</div>
				</div>
			</div>
			</div>

			
			<div class="modal fade" id="imageModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
			<div class="modal-dialog">
				<div class="modal-content">
				<div class="modal-header">
					<h5 class="modal-title" id="exampleModalLabel">Setup new Avatar</h5>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				<div class="modal-body">
					<form>
					
						<div class="input-group mb-3">
						<input type="file" @change="handleImageAvatar" accept="image/*" class="form-control" id="inputGroupFile02">
						
						</div>
					</form>
				</div>
				<div class="modal-footer">
					<button id="closeModalAvatarUpdate" type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
					<button @click="updateAvatar()" type="button" class="btn btn-warning">Save updates</button>
				</div>
				</div>
			</div>
			</div>
			



			</div>
			
			
			
			<div class="bottom">
			<span @click="goToFollowing()" class="following">
			<span class="count">{{profile.quantitySubscriptions}}</span>
			following
			</span>
			
			<span @click="goToFollowers()" class="followers">
				<span class="count">{{profile.quantitySubscribers}}</span>
				followers
			</span>
			<span @click="goToBanned()" class="banned" v-if="profile.me">
				<div @onclick="getMyProfile">
				banned</div>
			</span>
			
			
			
			</div>
			<div v-if="!profile.me">
			<button @click="subscribe" v-if="!this.profile.currentFollow" style="margin-top:5px" type="button" class="btn btn-warning">follow</button>
			<button @click="unsubscribe" v-else style="margin-top:5px" type="button" class="btn btn-warning">unfollow</button>
			<button @click="ban" v-if="!this.profile.currentBan" style="margin-top:5px; margin-left:10px" type="button" class="btn btn-danger">ban</button>
			<button @click="unban" v-else style="margin-top:5px; margin-left:10px" type="button" class="btn btn-danger">unban</button>
			</div>
			
		</div>
		
		</div>
		<div v-else>
			<LoadingSpinner>
        </LoadingSpinner>
		</div>

		



		
		
		<section id="gallery">
		<div class="container">
		<div v-if="!loadingPosts">
		<div v-if="postsProfile.length!=0" class="row">
		<PostProfileItem class="col-lg-4 mb-4" v-for="(post, index) in postsProfile" :key="index"
		:post="post"
		:profile="profile"
		@sendToComments="sendToComments(post)"
		@deletePostLike="deletePostLike(post)"
		@deletePostDislike="deletePostDislike(post)"
		@likePost="likePost(post)"
		@dislikePost="dislikePost(post)"
		@choosePost="choosePost(post)"

		>
		
			
		</PostProfileItem>
			
		
			
		</div>
		
		<div v-else>
			There is no data

		</div>

		</div>
		<div v-else>
		<LoadingSpinner>
        </LoadingSpinner>


		</div>
		
		</div>
		
		
		</section>
		
		
		
		


		<!--
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">MyProfile</h1>
			<div class="btn-toolbar mb-2 mb-md-0"
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
			</div>
		</div>
        <div>
            <p>This is some(myprofile page) component inside for checking data and then implement all necessary stuff</p>
        </div>-->
		<div v-if="profile.me" class="fab-container">
        
            
            <i data-bs-toggle="modal" data-bs-target="#postModal" style="font-size:4em;" class="bi bi-pencil-square"></i>
            
        
        
        </div>

		
	</div>
	<div v-observe-visibility="visibilityChanged"></div>
	</div>
</template>

<style>
@import url('https://fonts.googleapis.com/css?family=Open+Sans:300,400,700');

@import url('https://fonts.googleapis.com/icon?family=Material+Icons');
</style>