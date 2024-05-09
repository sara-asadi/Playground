using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Game.AnimalsFight
{
    public class AppDatabase
    {
        public static Fighter[] GetFighters()
        {

            string[] lines = File.ReadAllLines(".\\AnimalsFight\\Database.txt");
            List<Fighter> fighters = new List<Fighter>();

            foreach (var line in lines)
            {
                var info = line.Split(' ');
                fighters.Add(new Fighter(info[0], Int32.Parse(info[1])));
            }

            return fighters.ToArray();
        }
    }
}
