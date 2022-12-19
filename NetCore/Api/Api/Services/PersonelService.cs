using Api.Data;
using Api.Models;

namespace Api.Services
{
    public class PersonelService
    {
        public void SavePersonel(PersonelModel personelModel)
        {
            try
            {
                DbInsert(personelModel);
                RabbitMQInsert(personelModel);

            }
            catch (Exception)
            {

                throw;
            }
        }

        private void RabbitMQInsert(PersonelModel personelModel)
        {
            RabbitMQService.AddQueue("Personel", personelModel);
        }

        private void DbInsert(PersonelModel personelModel)
        {
            using (DataContext dataContext=new DataContext())
            {
                dataContext.Add<PersonelModel>(personelModel);  
                dataContext.SaveChanges();
            }
        }

        public IEnumerable<PersonelModel> GetPersonelModels()
        {
            var retPersonel = new List<PersonelModel>();

            return retPersonel;

        }

    }
}
