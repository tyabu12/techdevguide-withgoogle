import unittest

from main import string_splosion

class TestSplosion(unittest.TestCase):
    def test_splosion(self):
      # (input, expected)
      pairs = [
        ('Code', 'CCoCodCode'),
        ('abc', 'aababc'),
        ('ab', 'aab'),
        ('x', 'x'),
        ('fade', 'ffafadfade'),
        ('There', 'TThTheTherThere'),
        ('Kitten', 'KKiKitKittKitteKitten'),
        ('Bye', 'BByBye'),
        ('Good', 'GGoGooGood'),
        ('Bad', 'BBaBad')
      ]
      for s, expected in pairs:
        actual = string_splosion(s)
        self.assertEqual(expected, actual)

if __name__ == "__main__":
    unittest.main()