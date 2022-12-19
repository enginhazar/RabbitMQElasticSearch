<template>
<div>
  <!-- <h1> Merhaba {{ad}} </h1>
  <button @click="changeAd('hazar')"> Değiştir</button> -->
 <Personel></Personel>
</div>
</template>

<script>

 import Personel from './components/Personel'


export default {
  name: 'App',
  components: {
    Personel
  }
  // },
  // data() {
  //   return {
  //     ad:"engin"
  //   }
  // },
  // methods:{
  //   changeAd(changeAd){
  //     this.ad=changeAd
  //   }
  // }
}
</script>

<style>

</style>
