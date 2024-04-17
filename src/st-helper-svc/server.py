from concurrent import futures

import grpc
from grpc_reflection.v1alpha import reflection

import helper_pb2
import helper_pb2_grpc
from financial_analysis import FinancialAnalysisService


def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))

    helper_pb2_grpc.add_STHelperServicer_to_server(FinancialAnalysisService(), server)

    service_names = (
        helper_pb2.DESCRIPTOR.services_by_name['STHelper'].full_name,
        reflection.SERVICE_NAME,
    )

    reflection.enable_server_reflection(service_names, server)

    server_port = '[::]:50053'
    server.add_insecure_port(server_port)
    print(f"Helper server listening on {server_port}")
    server.start()

    try:
        server.wait_for_termination()
    except KeyboardInterrupt:
        print("Shutting down the server...")
        server.stop(0)


if __name__ == '__main__':
    serve()