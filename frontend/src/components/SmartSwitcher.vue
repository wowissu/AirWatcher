<template>
  <div class="smart-switcher">
    <select name="" id="">
      <option value="" v-for="(site, index) in sites" :key="index">
        {{ site.SiteName }}
      </option>
    </select>
  </div>
</template>

<script>
  import { mapState } from 'vuex';

  export default {
    name: 'SmartSwitcher',
    mounted() {
      //
    },
    computed: {
      ...mapState({
        sites: state => state.sites.data
      })
    },
    methods: {

    },
  }
</script>

<style lang="scss" scoped>
  .smart-switcher {

  }
</style>