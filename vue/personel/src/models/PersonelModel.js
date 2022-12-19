export default class Personel {
   

    tcKimlikNo;
    ad ;
    soyad ;
    maas ;
    dogumTarih ;
    adres;

 
    constructor(p_sicilNo,p_tckimlikNo,p_ad,p_soyad,p_maas,p_dogumTarih,p_adres) 
    {
      this.sicilNo=p_sicilNo;
      this.tcKimlikNo=p_tckimlikNo;
      this.ad=p_ad;
      this.soyad=p_soyad;
      this.maas=p_maas;
      this.dogumTarih=p_dogumTarih;
      this.adres=p_adres;
    }
         
}