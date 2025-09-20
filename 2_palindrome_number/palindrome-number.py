class Solution:
     def isPalindrome(self, x: int) -> bool:
          if x < 0:
            return False
          given_int = x
          created_int = 0
          while x > 0:
            reminder = x % 10
            print(reminder)
            created_int = (created_int * 10) + reminder
            x = int(x / 10)
          return created_int == given_int

if __name__ == "__main__":
    examples = [121, -121, 10]
    for index, example in enumerate(examples):
        print(f"Example {index}\nInput: {example} \nOutput: {Solution().isPalindrome(example)}")
