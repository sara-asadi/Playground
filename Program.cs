using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Game
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("What game do you wanna play?".ToUpper());
            var command = "";
            do
            {
                Console.WriteLine("1.RockPaperScissor");
                Console.WriteLine("2.Palindrome");

                command = Console.ReadLine();

                switch (command)
                {
                    case "1":
                        RockPaperScissor.Play();
                        break;
                    case "2":
                        Palindrome.Play();
                        break;

                }

            } while (command != "x");


        }
    }
}
