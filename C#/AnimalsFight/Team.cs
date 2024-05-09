using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Game.AnimalsFight
{
    public class Team
    {
        public static readonly Fighter[] fighters = AppDatabase.GetFighters();
        static Random random = new Random();

        public Fighter First { get; set; }
        public Fighter Second { get; set; }

        public Team()
        {
            First = fighters[random.Next(fighters.Length)];
            Second = fighters[random.Next(fighters.Length)];
        }

        public int GetPower()
        {
            return First.Power + Second.Power;
        }

        public void Introduce()
        {
            First.Introduce();
            Second.Introduce();
        }
    }
}
