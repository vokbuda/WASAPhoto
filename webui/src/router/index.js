import {createRouter, createWebHistory} from 'vue-router'
import Welcome from '../views/Welcome.vue'
import FeedView from '../views/FeedView.vue'
import MyProfileView from '../views/MyProfileView.vue'
import SearchProfileView from '../views/SearchProfileView.vue'
import Register from '../views/Register.vue'
import Application from '../views/Application.vue'
import Comments from '../views/Comments.vue'

import isAuthenticated from '../auth/auth'

const router = createRouter({
	history: createWebHistory(),
	routes: [
		
		{
			path: '/dologin', 
			component:Register
			
	
	
	
	
		},
		{
			path:'/',
			component: Application,
			
			children:
				[
				{path: '/welcome', component: Welcome},
				{
					path: '/posts', component: FeedView,
				},
				{path: '/posts/:postid/comments',component:Comments},
				{path: '/profiles/:userid', component: MyProfileView},
				{path: '/profiles', component: SearchProfileView}
				// /posts/:postid/comments
			]
			

		}

		

		
		// {path: '/some/:id/link', component: Welcome},
	]
})


// this function below u need for authorization

router.beforeEach(async (to, from,next) => {
	
	
	
	if(!sessionStorage.getItem("token") && to.fullPath!=='/dologin') {

		
	  next({path:'/dologin'})
	  
	}
	if(to.fullPath!=='/dologin'&& !sessionStorage.getItem("token")){
		next({path:'/dologin'})
	}
	else{
		next()
	}
	
	
	
	
	
	
	
  })


export default router
