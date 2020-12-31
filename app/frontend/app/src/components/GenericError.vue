<template>
    <p v-if="errors.length">
        <b>Ops!</b>
        <ul>
        <li v-for="(error, index) in errors" :key="index" class="text-danger"> {{ error.message }}</li>
        </ul>
    </p>
</template>

<script>

import { mapState } from 'vuex'
export default {
  name: 'GenericError',
  computed: {
    ...mapState({
      errors: state => state.common.generic_errors
    })
  }
}

</script>
