<template>
    <p v-if="errors.length">
        <b>Ocorreram os seguintes erros:</b>
        <ul>
        <li v-for="(error, index) in errors" :key="index" class="text-danger"> {{ error }}</li>
        </ul>
    </p>
</template>

<script>

import { mapState } from 'vuex'
export default {
  name: 'ValidationError',
  computed: {
    ...mapState({
      errors: state => state.auth.errors
    })
  }
}

</script>
