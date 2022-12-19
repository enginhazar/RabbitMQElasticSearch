using Api.Models;
using Microsoft.EntityFrameworkCore;

namespace Api.Data
{
    public class DataContext:DbContext
    {
        public DbSet<PersonelModel> Employees { get; set; }
  

        protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        {
            
            optionsBuilder.UseInMemoryDatabase("exampleDatabase");
        }
    }
}
