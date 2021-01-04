<template>
  <div>Portifolio</div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'

export default {
  computed: {
    ...mapGetters(['isAuthenticated', 'user'])
  },
  methods: {
    ...mapActions(['logout']),
    ...mapActions({
      exit: 'logout'
    })
  }
}
</script>
