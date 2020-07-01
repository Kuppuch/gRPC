using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Net.Sockets;
using System.Threading;
using System.Threading.Tasks;
using Microsoft.AspNetCore.Hosting;
using Microsoft.Extensions.Hosting;

namespace GrpcService2._0 {
    class Program {
        static Random rand = new Random();
        List<IObserver> observers = new List<IObserver>();
        static int number = 0;

        public static void Main(string[] args) {

            GreeterService gs = new GreeterService();
            CreateHostBuilder(args).Build().Start();
            
            
            
        }

        // Additional configuration is required to successfully run gRPC on macOS.
        // For instructions on how to configure Kestrel and gRPC clients on macOS, visit https://go.microsoft.com/fwlink/?linkid=2099682
        public static IHostBuilder CreateHostBuilder(string[] args) =>
            Host.CreateDefaultBuilder(args)
                .ConfigureWebHostDefaults(webBuilder => {
                    webBuilder.UseStartup<Startup>();
                });

        



        //public static void Common() {
        //    while (true) {
        //        number = rand.Next(0, 20);
        //        //Console.WriteLine("___ " + number);
        //        NotifyObservers();
        //        Thread.Sleep(500 + rand.Next(0, 500));
                
        //    }
        //}

    }


}
