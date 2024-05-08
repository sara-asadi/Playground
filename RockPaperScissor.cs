using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Game
{
    public class RockPaperScissor
    {
        private static Dictionary<int, string> Choices = new Dictionary<int, string> { { 1, "Rock" }, { 2, "Paper" }, { 3, "Scissor" } };

        public static void Play()
        {
            int compScore = 0, userScore = 0, winningScore = 3;

            int userChoice;

            Console.WriteLine("Game Started".ToUpper());

            while (compScore < winningScore && userScore < winningScore)
            {
                Console.WriteLine("Enter Your Choice");
                Console.WriteLine("1:Rock");
                Console.WriteLine("2:Paper");
                Console.WriteLine("3:Scissor");

                userChoice = Int32.Parse(Console.ReadLine());

                if (userChoice != 1 && userChoice != 2 && userChoice != 3)
                {
                    Console.WriteLine("You made the wrong Choice!");
                    continue;
                }

                int randomNumber = new Random().Next(1, 3);

                int result = CheckWinner(userChoice, randomNumber);

                if (result == 0)
                {
                    Console.WriteLine("Draw!".ToUpper());
                }
                else if (result > 0)
                {
                    userScore++;
                    Console.WriteLine("You won!".ToUpper());
                }
                else
                {
                    compScore++;
                    Console.WriteLine("You Lost!".ToUpper());
                }
            }


            if (userScore > compScore)
            {
                Console.WriteLine("You Won The Game!");
            }
            else
            {
                Console.WriteLine("You Lost the Game!");
            }
        }

        private static int CheckWinner(int userChoice, int compChoice)
        {
            Console.WriteLine($"Computer Chose {Choices[compChoice]}");

            if (userChoice == compChoice)
            {
                return 0;
            }

            switch (userChoice)
            {
                case 1:
                    return compChoice == 2 ? -1 : 1;
                case 2:
                    return compChoice == 3 ? -1 : 1;
                case 3:
                    return compChoice == 1 ? -1 : 1;
                default:
                    throw new ArgumentException();
            }
        }

    }
}
