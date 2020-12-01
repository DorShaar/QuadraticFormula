import unittest
from EquationsReader import EquationsReader
EquationsReader = EquationsReader.EquationsReader

class EquationReaderTest(unittest.TestCase):

    def test_getQueueMessages_csvContainsThreeGoodAndOneBadEquations_get3GoodEquations(self):
        csv_path = 'test_files/Equations.csv'

        equationsReader = EquationsReader(csv_path)
        queue_messages = equationsReader.get_queue_messages()

        expectedSize = 3
        self.assertEqual(
            expectedSize,
            len(queue_messages),
            msg="Queue messages size is expected to be 3")

if __name__ == '__main__':
    unittest.main()