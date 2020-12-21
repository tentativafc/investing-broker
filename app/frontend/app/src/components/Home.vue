<template>
    <div>
        <div class="container">
            <div class="row" v-if="isAuthenticated" >
                Logged Ared
            </div>  
            <div class="row" v-if="!isAuthenticated">
                <div class="col-6">       
                    <a href="/login" class="btn-lg">Login</a>
                </div>
                <div class="col-6">  
                    <a href="/register" class="btn-lg">Register</a> 
                </div>
            </div>  
        </div>  
    </div>
</template>

<script>

// import { mapGetters } from "vuex";

// export default {
//     computed: {
//         ...mapGetters(["isAuthenticated"])
//     }
// }
</script>