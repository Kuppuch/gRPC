﻿using Grpc.Net.Client;
using GrpcService1;
using System;
using System.Threading.Tasks;

namespace GRPCGreeterClient {
    class Program {
        //static void Main(string[] args) {
        //    Console.WriteLine("Hello World!");
        //}

        static async Task Main(string[] args) {
            // The port number(5001) must match the port of the gRPC server.
            using var channel = GrpcChannel.ForAddress("https://localhost:5001");
            var client = new Greeter.GreeterClient(channel);
            //var reply = await client.SayHelloAsync(
                              //new HelloRequest { Name = "Kuppuch" });
            //Console.WriteLine("Greeting: " + reply.Message);
            var reply = await client.GetRandNumAsync(
                  new NumRequest { Name = "Kuppuch" });
            Console.WriteLine("Greeting: " + reply.Message);
            Console.WriteLine("Press any key to exit...");
            Console.ReadKey();
        }
    }
}
