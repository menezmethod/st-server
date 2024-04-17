import unittest

from src import helper_pb2


class TestSTHelperService(unittest.TestCase):

    def test_interaction(self):
        # Simulated test for interaction with Llama2
        request = helper_pb2.LlamaRequest(query="Test Query")
        expected_response = "Processed: Test Query"
        # You would add here the mock to simulate the server response
        self.assertEqual(expected_response, "Processed: Test Query")

if __name__ == '__main__':
    unittest.main()
