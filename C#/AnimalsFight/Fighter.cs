using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Game.AnimalsFight
{
    public class Fighter
    {
        string _name;
        int _power;

        public int Power { get => _power; }

        public Fighter(string name, int power)
        {
            _name = name;
            _power = power;

        }

        public void Introduce()
        {
            Console.WriteLine($"Hi I am {_name}");
        }
    }
}
