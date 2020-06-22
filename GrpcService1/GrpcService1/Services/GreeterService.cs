using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;
using Grpc.Core;
using Microsoft.Extensions.Logging;

namespace GrpcService1 {
    public class GreeterService : Greeter.GreeterBase {
        private readonly ILogger<GreeterService> _logger;
        public GreeterService(ILogger<GreeterService> logger) {
            _logger = logger;
        }
        Random rand = new Random();

        public override Task<HelloReply> SayHello(HelloRequest request, ServerCallContext context) {
            return Task.FromResult(new HelloReply {
                Message = "Hello " + request.Name
            });
        }

        /*public override Task<NumReply> GetRandNum(NumRequest request, ServerCallContext context) {
            return Task.FromResult(new NumReply {
                Message = Convert.ToString(rand.Next())
            });
        }*/

        public override async Task GetRandNum(NumRequest request, IServerStreamWriter<NumReply> responseStream, ServerCallContext context) {
           
            for (int i = 0; i < 10; i++) {
                var resp = new NumReply {
                    Message = Convert.ToString(rand.Next(1000))
                };
                await responseStream.WriteAsync(resp);
                Thread.Sleep(1000);
            }
        }
    }
}
