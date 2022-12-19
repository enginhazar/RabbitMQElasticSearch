using Api.Models;
using Api.Services;
using Microsoft.AspNetCore.Http;
using Microsoft.AspNetCore.Mvc;

namespace Api.Controller
{
    [Route("api/[controller]")]
    [ApiController]
    public class PersonelController : ControllerBase
    {
        private PersonelService  personelService;
        public PersonelController(PersonelService _personelService) {
            personelService=_personelService;
        }

        [HttpGet]
        public string Get()
        {
            return "engin";
        }

        [HttpPost]
        public IActionResult SavePersonel(PersonelModel personelModel)
        {
            personelService.SavePersonel(personelModel);

            return Ok();
        }
    }
}
