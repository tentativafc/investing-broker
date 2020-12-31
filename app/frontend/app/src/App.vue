<template>
  <div>
    <GenericError />
    <router-view></router-view>
  </div>
</template>

<script>
import GenericError from '@/components/GenericError'
export default {
  name: 'App',
  components: { GenericError }
}
</script>
