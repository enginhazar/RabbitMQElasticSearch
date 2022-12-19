using System.ComponentModel.DataAnnotations;

namespace Api.Models
{
    public class PersonelModel
    {
       
        public int SicilNo { get; set; }
        [Key]
        public Int64 TcKimlikNo { get; set; }
        public string Ad { get; set; }
        public string Soyad { get; set; }

        public double Maas { get; set; }
        public DateTime DogumTarih { get; set; }
        public string Adres { get; set; }
    }
}
