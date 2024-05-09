using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;

namespace Game
{
    public class Palindrome
    {
        public static void Play()
        {
            var command = "";
            do
            {
                Console.WriteLine("Enter a Word:".ToUpper());
                command = Console.ReadLine();
                var result = CheckResult(command);
                if (result)
                {
                    Console.WriteLine("Succeed");
                }
                else
                {
                    Console.WriteLine("Failed");
                }
            } while (command != "x");
        }

        public static bool CheckResult(string word)
        {
            var wordReverse = "";

            for (int i = word.Length - 1; i >= 0; i--)
            {
                wordReverse += word[i];
            }

            if (word == wordReverse)
            {
                return true;
            }
            else
            {
                return false;
            }
        }
    }
}
