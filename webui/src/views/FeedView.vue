<script>
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
		this.refresh()
	}
}

</script>
<style scoped>
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
			<img src="https://images.unsplash.com/photo-1516214104703-d870798883c5?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=700&q=60" alt="" class="card-img-top">
			<div class="card-body">
				<img class="card-user avatar avatar-large" src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fsafeharborpartners.com%2Fwp-content%2Fuploads%2Fshutterstock_169562684-449x375.jpg&f=1&nofb=1&ipt=fe4b42d35bb3eb2cf3d88d1eb7ebcb7e883e15736e51a2db2367cbf4f9eca201&ipo=images">

				<h5 class="card-title">{{post.username}}</h5>
				<p class="card-text">{{post.text}}</p>
				<div class="in_a_row">
				<btn class="btn btn-outline-success btn-sm">Comments</btn>

				<btn class="btn btn-outline-danger btn-sm"><i class="bi bi-hand-thumbs-up"></i></btn>
				<h5>32k</h5>

				<btn href="" class="btn btn-outline-danger btn-sm"><i class="bi bi-hand-thumbs-down"></i></btn>
				<h5>64k</h5>
				</div>
				
			</div>
			</div>
			</div>
		<div class="col-lg-4 mb-4">
		<div class="card">
			<img src="https://images.unsplash.com/photo-1516214104703-d870798883c5?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=700&q=60" alt="" class="card-img-top">
			<div class="card-body">
				<img class="card-user avatar avatar-large" src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fsafeharborpartners.com%2Fwp-content%2Fuploads%2Fshutterstock_169562684-449x375.jpg&f=1&nofb=1&ipt=fe4b42d35bb3eb2cf3d88d1eb7ebcb7e883e15736e51a2db2367cbf4f9eca201&ipo=images">

				<h5 class="card-title">Sunset</h5>
				<p class="card-text">Lorem ipsum dolor sit amet consectetur, adipisicing elit. Ut eum similique repellat a laborum, rerum voluptates ipsam eos quo tempore iusto dolore modi dolorum in pariatur. Incidunt repellendus praesentium quae!</p>
				<div class="in_a_row">
				<btn class="btn btn-outline-success btn-sm">Comments</btn>

				<btn class="btn btn-outline-danger btn-sm"><i class="bi bi-hand-thumbs-up"></i></btn>
				<h5>32k</h5>

				<btn href="" class="btn btn-outline-danger btn-sm"><i class="bi bi-hand-thumbs-down"></i></btn>
				<h5>64k</h5>
				</div>
				
			</div>
			</div>
			</div>
		<div class="col-lg-4 mb-4">
		<div class="card">
			<img src="https://images.unsplash.com/photo-1516214104703-d870798883c5?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=crop&w=700&q=60" alt="" class="card-img-top">
			<div class="card-body">
				<img class="card-user avatar avatar-large" src="https://external-content.duckduckgo.com/iu/?u=https%3A%2F%2Fsafeharborpartners.com%2Fwp-content%2Fuploads%2Fshutterstock_169562684-449x375.jpg&f=1&nofb=1&ipt=fe4b42d35bb3eb2cf3d88d1eb7ebcb7e883e15736e51a2db2367cbf4f9eca201&ipo=images">

				<h5 class="card-title">Sunset</h5>
				<p class="card-text">Lorem ipsum dolor sit amet consectetur, adipisicing elit. Ut eum similique repellat a laborum, rerum voluptates ipsam eos quo tempore iusto dolore modi dolorum in pariatur. Incidunt repellendus praesentium quae!</p>
				<div class="in_a_row">
				<btn href="" class="btn btn-outline-success btn-sm">Comments</btn>

				<btn class="btn btn-outline-danger btn-sm"><i class="bi bi-hand-thumbs-up"></i></btn>
				<h5>32k</h5>

				<btn href="" class="btn btn-outline-danger btn-sm"><i class="bi bi-hand-thumbs-down"></i></btn>
				<h5>64k</h5>
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
