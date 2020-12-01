from EquationsReader import EquationsReader

if __name__ == "__main__":
    csv_path = '../Resources/Equations.csv'

    equationsReader = EquationsReader.EquationsReader(csv_path)
    queue_messages = equationsReader.get_queue_messages()

    for message in queue_messages:
        equationsReader.send_queue_message(message)
