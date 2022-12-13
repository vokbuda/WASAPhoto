import {createRouter, createWebHashHistory} from 'vue-router'
import Welcome from '../views/Welcome.vue'
import FeedView from '../views/FeedView.vue'
import MyProfileView from '../views/MyProfileView.vue'
import SearchProfileView from '../views/SearchProfileView.vue'
import Register from '../views/Register.vue'
import Application from '../views/Application.vue'

const router = createRouter({
	history: createWebHashHistory(import.meta.env.BASE_URL),
	routes: [
		
		{
			path: '/register', 
			component:Register
			
	
	
	
	
		},
		{
			path:'/application',
			component: Application,
			children:[
				[{path: '/welcome', component: Welcome},
				{path: '/posts', component: FeedView},
				{path: '/profiles/:userid', component: MyProfileView},
				{path: '/profiles', component: SearchProfileView}]
			]

		}

		

		
		// {path: '/some/:id/link', component: Welcome},
	]
})


export default router
