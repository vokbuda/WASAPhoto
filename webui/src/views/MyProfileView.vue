<script>
import router from '../router'
import Post from '../entities/Post'
export default {
	data: function() {
		return {
			errormsg: null,
			loading: false,
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
			postsProfile:[]
		}
	},
	methods: {
		async refresh() {
			this.loading = true;
			this.errormsg = null;

			
		},
		handleImage(e){
			
			const selected_image=e.target.files[0]
			this.createBase64Image(selected_image)

		},
		likePost(post){
			this.userid=sessionStorage.getItem("userid")
			const postData = JSON.stringify({ "idPostEmotion": post.postid,"idUser":parseInt(this.userid) });
			
			
			
			this.$axios.put('s/posts/'+post.postid+'/like/'+this.userid, postData,{
				headers:{
					"Authorization":'Bearer '+sessionStorage.getItem("token")
				}
			}
			)
			.then((response)=> {
				
				if (response.status==204){
					
					//document.getElementById("closeModalPostCreate").click()
					//this.createdPost()
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
			
			console.log(sessionStorage.getItem('token'))
			__API_URL__
			
			fetch(__API_URL__+'/posts/'+post.postid+'/like/'+this.userid,{
				method:'DELETE',
				headers:{
					"Authorization":'Bearer '+sessionStorage.getItem("token")
				},
				body:postData
			}
			)
			.then((response)=> {
				
				if (response.status==204){
					post.currentemotion=0
					//document.getElementById("closeModalPostCreate").click()
					//this.createdPost()
					
				}
			})
			.catch(function (error) {
				console.log(error);
			});

		},
		dislikePost(post){
			this.userid=sessionStorage.getItem("userid")
			const postData = JSON.stringify({ "idPostEmotion": post.postid,"idUser":parseInt(this.userid) });
			
			
			
			this.$axios.put('/posts/'+post.postid+'/dislike/'+this.userid, postData,{
				headers:{
					"Authorization":'Bearer '+sessionStorage.getItem("token")
				}
			}
			)
			.then((response)=> {
				if (response.status==200){
					
					//document.getElementById("closeModalPostCreate").click()
					//this.createdPost()
					console.log("success response from database")
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
					"Authorization":'Bearer '+sessionStorage.getItem("token")
				},
				body:postData
			}
			)
			.then((response)=> {
				
				if (response.status==204){
					post.currentemotion=0
					//document.getElementById("closeModalPostCreate").click()
					//this.createdPost()
					
				}
			})
			.catch(function (error) {
				console.log(error);
			});

		},
		createdPost(){
			document.getElementById("liveToast").style.display="block";
			//setTimeout(1500)
			const delay = ms => new Promise(res => setTimeout(res, ms));


			const yourFunction = async () => {
				await delay(1500);
				document.getElementById("liveToast").style.display="none";
			};
			yourFunction()
						
			

		},
		createPost(){
			
			const postData = JSON.stringify({ "text": this.postText,"image":this.postImage });
			
			
			this.$axios.post('/profiles/'+this.$route.params.userid+'/posts', postData,{
				headers:{
					"Authorization":'Bearer '+sessionStorage.getItem("token")
				}
			}
			)
			.then((response)=> {
				if (response.status==200){
					console.log(response.data.postid)
					const post=new Post(this.$route.params.userid,response.data.postid,this.postText,
					'0','0',true,this.postImage,'',0)
					this.postsProfile.unshift(post)
					
					document.getElementById("closeModalPostCreate").click()
					this.createdPost()
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
		async getMyPosts(){
			
			this.header="Bearer "+sessionStorage.getItem("token")
		
		
			//this.userid=sessionStorage.getItem("userid")


			
		
			await this.$axios.get('/profiles/'+this.$route.params.userid+'/posts?offset=0',{
				headers:{
					'Authorization':this.header,
				
				}




			})
				.then(response => {
					// Handle response
					
					this.postsProfile=response.data
					console.log(this.postsProfile)
				})
				.catch(err => {
					// Handle errors
					console.error(err);
				
				});
					


			// then here u should setup your posts of your profile

		},
		sendToComments(postid){
			router.push('/posts/'+postid+'/comments')
			

		},
		async getMyProfile(){
		
		
			this.header="Bearer "+sessionStorage.getItem("token")
		
		
			//this.userid=sessionStorage.getItem("userid")


			
		
			await this.$axios.get('/profiles/'+this.$route.params.userid,{
				headers:{
					'Authorization':this.header,
				
				}




			})
				.then(response => {
					// Handle response
					
					this.profile=response.data
				})
				.catch(err => {
					// Handle errors
					console.error(err);
				
				});
					}
	},
	mounted() {
		this.postCreation=true
		this.getMyProfile()
		this.getMyPosts()
		
	}
}
</script>

<style scoped>
body{
  min-height:100vh;
  min-width:100vw;
  background:linear-gradient(to right bottom,rgba(blue,0.5),rgba(green,0.5)),url(http://i.imgur.com/woUbg3p.jpg);
  background-size:cover;
  background-position:center center;
  display:flex;
  justify-content:center;
  align-items:center;
}

#header{
  width:50vw;
  max-width:400px;
  min-width:300px;
  height:70px;
  background:url(http://i.imgur.com/woUbg3p.jpg);
  background-size:cover;
  background-position: center center;
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
    
    border:3px solid white;
    
    height:80px;
    width:100px;
    border-radius:10px;
    
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
.fab-container {
  
  position: fixed;
  
	z-index: 1;
  bottom: 30px; right: 30px;
  
}

</style>

<template>
	<div>

<div id="looker" class="position-fixed bottom-0 end-0 p-3" style="z-index: 11">
  <div id="liveToast" class="toast hide" role="alert" aria-live="assertive" aria-atomic="true">
    <div class="toast-header">
      
      <strong class="me-auto">WASAPhoto</strong>
      
      
    </div>
    <div class="toast-body">
      Post had been created
    </div>
  </div>
</div>
		<div class="shadow overflow">
		<div id="header"></div>
		<div id="profile">
			<div class="absolution">
			<div id="element1">
			
			<div class="image" v-if="this.profile.avatar"><img :src="person.avatar"></div>
			<div class="image" v-else><img src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fsafeharborpartners.com%2Fwp-content%2Fuploads%2Fshutterstock_169562684-449x375.jpg&f=1&nofb=1&ipt=fe4b42d35bb3eb2cf3d88d1eb7ebcb7e883e15736e51a2db2367cbf4f9eca201&ipo=images"></div>
			<!--<img src="https://a4-images.myspacecdn.com/images03/2/85a286a4bbe84b56a6d57b1e5bd03ef4/300x300.jpg" alt="" />-->
			
			</div>
			<div id="element2">
			<btn data-bs-toggle="modal" data-bs-target="#imageModal"><i class="bi bi-upload"></i></btn>
			</div>
			</div>
			<div class="name">
			{{this.profile.username}}
			<btn data-bs-toggle="modal" data-bs-target="#usernameModal"><i class="bi bi-pencil-square"></i></btn>
			Password<btn data-bs-toggle="modal" data-bs-target="#passwordModal"><i class="bi bi-pencil-square"></i></btn>
			
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
						<label for="recipient-name" class="col-form-label">Current password:</label>
						<input type="text" class="form-control" id="recipient-name">
					</div>
					<div class="mb-3">
						<label for="message-text" class="col-form-label">New username:</label>
						<input type="text" class="form-control" id="message-text">
					</div>
					</form>
				</div>
				<div class="modal-footer">
					<button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
					<button type="button" class="btn btn-primary">Save updates</button>
				</div>
				</div>
			</div>
			</div>
			<div v-if="this.postCreation" class="modal fade" id="postModal" tabindex="-1" aria-labelledby="exampleModalLabel" >
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
						<input type="file" @change="this.handleImage" accept="image/*" class="form-control" id="inputGroupFile02">
						
						</div>
					</div>
					<div class="mb-3">
						<label for="message-text" class="col-form-label">Text:</label>
						<input v-model="this.postText" type="text" class="form-control" id="message-text">
					</div>
					</form>
				</div>
				<div class="modal-footer">
					<button id="closeModalPostCreate" type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
					<button type="button" @click="this.createPost()" class="btn btn-primary">Create</button>
				</div>
				</div>
			</div>
			</div>
			<div class="modal fade" id="passwordModal" tabindex="-1" aria-labelledby="exampleModalLabel" aria-hidden="true">
			<div class="modal-dialog">
				<div class="modal-content">
				<div class="modal-header">
					<h5 class="modal-title" id="exampleModalLabel">Setup new Password</h5>
					<button type="button" class="btn-close" data-bs-dismiss="modal" aria-label="Close"></button>
				</div>
				<div class="modal-body">
					<form>
					<div class="mb-3">
						<label for="oldpassword" class="col-form-label">Current password:</label>
						<input type="password" class="form-control" id="oldpassword">
					</div>
					<div class="mb-3">
						<label for="newpassword" class="col-form-label">New password:</label>
						<input type="password" class="form-control" id="newpassword">
					</div>
					</form>
				</div>
				<div class="modal-footer">
					<button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
					<button type="button" class="btn btn-primary">Save updates</button>
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
					<div class="mb-3">
						<label for="oldpassword" class="col-form-label">Current password:</label>
						<input type="password" class="form-control" id="oldpassword">
					</div>
						<div class="input-group mb-3">
						<input type="file" class="form-control" id="inputGroupFile02">
						
						</div>
					</form>
				</div>
				<div class="modal-footer">
					<button type="button" class="btn btn-secondary" data-bs-dismiss="modal">Close</button>
					<button type="button" class="btn btn-primary">Save updates</button>
				</div>
				</div>
			</div>
			</div>
			<!--<div class="input-group mb-3">
  <input type="file" class="form-control" id="inputGroupFile02">
  <label class="input-group-text" for="inputGroupFile02">Upload</label>
</div>-->



			</div>
			
			
			<!--
			<div class="nickname">
			@daftpunk
			</div>
			<div class="location">
			<i class="material-icons">place</i>Europe
			</div>
			-->
			<div class="bottom">
			<span class="following">
			<span class="count">{{this.profile.quantitySubscriptions}}</span>
			following
			</span>
			
			<span class="followers">
				<span class="count">{{this.profile.quantitySubscribers}}</span>
				followers
			</span>
			<span class="banned" v-if="this.profile.me">
				<div @onclick="this.getMyProfile">
				banned</div>
			</span>
			
			
			
			</div>
			<div v-if="!this.profile.me">
			<button v-if="this.switcherSubscription" style="margin-top:5px" type="button" class="btn btn-warning">{{this.subscriptionText}}</button>
			<button v-if="this.switcherBan" style="margin-top:5px; margin-left:10px" type="button" class="btn btn-danger">{{this.banText}}</button>
			</div>
			
		</div>
		
		</div>
		
		
		<section id="gallery">
		<div class="container">
			<div class="row">
			<div class="col-lg-4 mb-4" v-for="(post, index) in postsProfile" :key="index">
		<div class="card">
			<img v-bind:src="'data:image/jpeg;base64,'+post.image" alt="" class="card-img-top">
			<div class="card-body">
				<img class="card-user avatar avatar-large" src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fsafeharborpartners.com%2Fwp-content%2Fuploads%2Fshutterstock_169562684-449x375.jpg&f=1&nofb=1&ipt=fe4b42d35bb3eb2cf3d88d1eb7ebcb7e883e15736e51a2db2367cbf4f9eca201&ipo=images">

				<h5 class="card-title">{{this.profile.username}}</h5>
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
		<div
			class="d-flex justify-content-between flex-wrap flex-md-nowrap align-items-center pt-3 pb-2 mb-3 border-bottom">
			<h1 class="h2">MyProfile</h1>
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
			</div>
		</div>
        <div>
            <p>This is some(myprofile page) component inside for checking data and then implement all necessary stuff</p>
        </div>-->
		<div v-if="this.profile.me" class="fab-container">
        
            
            <i data-bs-toggle="modal" data-bs-target="#postModal" style="font-size:4em;" class="bi bi-pencil-square"></i>
            
        
        
        </div>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
@import url('https://fonts.googleapis.com/css?family=Open+Sans:300,400,700');

@import url('https://fonts.googleapis.com/icon?family=Material+Icons');
</style>