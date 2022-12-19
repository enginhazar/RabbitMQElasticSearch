import Personel from "@/models/PersonelModel"

const  url='http://localhost:5013';

export default {


    getPersonel(){
        const personel=new Personel(123,1234567890,"engin","hazar",141.47,'2022-12-29',"yüzüncü yıl mahallesi");
        return personel; 
    },

   async savePersonel(personelModel ){
        
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(personelModel)
        };
         await fetch(url+"/api/Personel", requestOptions) .then(async response => {
            //const data = await response;
      
            // check for error response
            if (!response.ok) {
                const data = await response;
              // get error message from body or default to response status
             const error = (data && data.message) || response.status;
              return Promise.reject(error);
            }
            //onst data = await response;
         
          })
          .catch(error => {
            console.log(error);
            console.error("error yaz");
            this.errorMessage = error;
            console.error('There was an error!', error);
          });
       

    }
}