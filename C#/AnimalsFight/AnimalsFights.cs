using Game.AnimalsFight;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Game
{
    public static class AnimalsFights
    {
        public static void Play()
        {
            while (true)
            {

                var team1 = new Team();
                var team2 = new Team();

                Console.WriteLine("Team1:");
                team1.Introduce();
                Console.WriteLine("Team2:");
                team2.Introduce();

                Console.WriteLine("Who win?");

                string userChoice = Console.ReadLine();
                if(userChoice == "x")
                {
                    break;
                }

                if (team1.GetPower() > team2.GetPower())
                {
                    Console.WriteLine(userChoice == "1" ? "You were correct! Team1 Won :)" : "You were wrong! Team2 Won :(");
                }
                else if (team1.GetPower() < team2.GetPower())
                {
                    Console.WriteLine(userChoice == "2" ? "You were correct! Team2 Won :)" : "You were wrong! Team1 Won :(");
                }
                else
                {
                    Console.WriteLine("It was a Draw! No one Won the Fight :|");
                }
            }


        }

    }
}
