using Grpc.Core;
using Grpc.Net.Client;
using GrpcService1;
using System;
using System.Threading;
using System.Threading.Tasks;

namespace GRPCGreeterClient {
    class Program {

        static async Task Main(string[] args) {
            // The port number(5001) must match the port of the gRPC server.
            using var channel = GrpcChannel.ForAddress("https://localhost:5001");
            var client = new Greeter.GreeterClient(channel);

            //AppContext.SetSwitch("System.Net.Http.SocketsHttpHandler.Http2UnencryptedSupport", true);
            //var channel = GrpcChannel.ForAddress("http://localhost:5001");
            //var client = new Greeter.GreeterClient(channel);


            /*var reply = await client.GetRandNumAsync(
                  new NumRequest { Name = "Kuppuch" });
            Console.WriteLine("Greeting: " + reply.Message);*/

            using var call = client.GetRandNum(new NumRequest { Name = "World" });

            await foreach (var response in call.ResponseStream.ReadAllAsync()) {
                Console.WriteLine(response.Message);
                Thread.Sleep(5000);
            }

            Console.WriteLine("Press any key to exit...");
            Console.ReadKey();
        }
    }
}
