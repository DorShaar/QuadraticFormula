from EquationsReader import EquationsReader

if __name__ == "__main__":
    csv_path = '../Resources/Equations.csv'

    equationsReader = EquationsReader(csv_path)
    equationsReader.print_queue_messages()