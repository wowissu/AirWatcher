<template>
  <div class="rwd-container">
    <slot></slot>
  </div>
</template>

<script>
  export default {
    name: 'RWDContainer'
  }
</script>

<style lang="scss" scoped>
  .rwd-container {
    width: 100vw;
    margin: 0 auto;
    min-height: 100vh;
    max-width: 480px;

    @include tablet {
      max-width: 768px;
    }

    @include desktop {
      max-width: 1024px;
    }
  }
</style>