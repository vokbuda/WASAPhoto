<script>
import {adjustNumber} from '../helpers/adjustNumber.js'
export default {
  props: ['person'],
 
    methods:{
        adjustNumber(data){
            return adjustNumber(data)
        },
        gotoProfile(){
            this.$emit("gotoProfile")
        },
        ban(){
            this.$emit("ban")
        },
        unban(){
            this.$emit("unban")
        },
        unsubscribe(){
            this.$emit("unsubscribe")
        }
        


    },
    mounted(){
        console.log("Below u can check some data for current component")
        console.log(this.post)
        
    }    
}
</script>
<style scoped src="@/assets/followingUser.css">


</style>

<template>
<div>

  <div  class="profile-card-4 text-center">
						<div @click="gotoProfile(person.userid)" v-if="person.avatar"><img :src="'data:image/jpeg;base64,'+person.avatar" class="card-img-top"></div>
						<div @click="gotoProfile(person.userid)" v-else><img src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fsafeharborpartners.com%2Fwp-content%2Fuploads%2Fshutterstock_169562684-449x375.jpg&f=1&nofb=1&ipt=fe4b42d35bb3eb2cf3d88d1eb7ebcb7e883e15736e51a2db2367cbf4f9eca201&ipo=images" class="card-img-top"></div>
						<div class="profile-content">
							<div class="profile-name" @click="gotoProfile(person.userid)">{{person.username}}</div>
						</div>
                        <div v-if="person.mine" class="row">
								
								<div class="col-xs-4">
									<button @click="unsubscribe(person)" style="margin-bottom:10px" type="button" class="btn btn-warning">unfollow</button>
								</div>
                                <div v-if="!person.currentBan" class="col-xs-4">
                                    <button @click="ban(person)" style="margin-bottom:10px" type="button" class="btn btn-danger">ban</button>
                                    
									
								</div>
                                <div v-else class="col-xs-4">
                                    <button @click="unban(person)" style="margin-bottom:10px" type="button" class="btn btn-danger">unban</button>
                                    
									
								</div>
								
							</div>
					</div>



        


        


</div>
</template>                   