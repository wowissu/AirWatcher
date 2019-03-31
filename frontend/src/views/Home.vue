<template>
  <div class="home">
    <RWDContainer>
      <SmartSwitcher></SmartSwitcher>
    </RWDContainer>
  </div>
</template>

<script>
import RWDContainer from '@/components/RWDContainer'
import SmartSwitcher from '@/components/SmartSwitcher'

export default {
  name: 'home',
  components: {
    RWDContainer,
    SmartSwitcher
  }
}
</script>

<style lang="scss">
.home {

}
</style>