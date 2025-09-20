class Solution:
    mapping = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
    def romanToInt(self, s: str) -> int:
        total = 0
        for i in range(len(s)):
            current_value = self.mapping[s[i]]
            if i + 1 < len(s) and self.mapping[s[i+1]] > current_value:
                total -= current_value
            else:
                total += current_value        
        return total

if __name__ == "__main__":
    examples = ["III", "LVIII", "MCMXCIV"]
    for index, example in enumerate(examples):
        print(f"Example {index}\nInput: {example} \nOutput: {Solution().romanToInt(example)}")
