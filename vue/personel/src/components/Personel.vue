<template>

    <div class="container">
        <div class="row mb-4">
            <h1>Personel Kartı</h1>
        </div>
        <div class="row mb-4">
    <div class="col">
      <div class="form-outline">
        <input type="number" v-model="personelModel.sicilNo" id="formSicilNo" class="form-control" />
        <label class="form-label" for="form3Example1">Sicil No</label>
      </div>
    </div>
    <div class="col">
      <div class="form-outline">
        <input type="text" id="formTcKimlik"  v-model="personelModel.tcKimlikNo" class="form-control" />
        <label class="form-label" for="form3Example2">TcKimlikNo</label>
      </div>
    </div>
  </div>
  <div class="row mb-4">
    <div class="col">
      <div class="form-outline">
        <input type="text" id="formAd"  v-model="personelModel.ad" class="form-control" />
        <label class="form-label" for="form3Example1">Adı</label>
      </div>
    </div>
    <div class="col">
      <div class="form-outline">
        <input type="text" id="formSoyad"  v-model="personelModel.soyad" class="form-control" />
        <label class="form-label" for="form3Example2">Soyadı</label>
      </div>
    </div>
  </div>
<div class="row mb4">
    <div class="col">
  <div class="form-outline mb-4">
    <input type="date" id="formMaas"  v-model="personelModel.dogumTarih" class="form-control" />
    <label class="form-label" for="form3Example3">Doğum Tarihi</label>
  </div>
</div>
<div class="col">
  <div class="form-outline mb-4">
    <input type="number" id="formMaas"  v-model="personelModel.maas" class="form-control" />
    <label class="form-label" for="form3Example3">Maaşı</label>
  </div>
</div>
<div class="form-outline mb-4">
    <textarea  type="text" id="formAdres"  v-model="personelModel.adres" class="form-control" />
    <label class="form-label" for="form3Example3">Adres</label>
  </div>
</div>
 


  <button  @click="personelSave" class="btn btn-primary btn-block mb-4">Kaydet</button>

</div>

</template>

<script>
import PersonelService from '@/services/PersonelService'

export default {
    name:"PersonelComponent",
    methods:{
        personelSave(){
          
            PersonelService.savePersonel(this.personelModel);
        }

    },

   data(){
    //const personel=PersonelService.getPersonel();
    return {
        personelModel:PersonelService.getPersonel()
    }
   }


}
</script>

<style>

</style>