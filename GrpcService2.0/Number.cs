using Grpc.Core;
using GrpcService1;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading;
using System.Threading.Tasks;

namespace GrpcService2._0{
    public class Number : IObservable {

        List<IObserver> observers = new List<IObserver>();
        List<IObserver> observers2Remove = new List<IObserver>();
        public int ValueNumber { get; set; }

        public Number() {
            Thread common = new Thread(new ThreadStart(Generate)) {
                Name = "Общий поток"
            };
            common.Start();
        }

        public void RegisterObserver(IObserver o) {
            observers.Add(o);
        }

        public void RemoveObserver(IObserver o) {
            observers2Remove.Add(o);
        }

        public void NotifyObservers() {
            foreach (IObserver o in observers2Remove) {
                observers.Remove(o);
            }
            observers2Remove.Clear();
            foreach (IObserver o in observers) {
                o.Update(new Notification() { observer = o, number = this});
                Console.WriteLine("Update запущен");
            }
        }

        public void Generate() {
            Random rand = new Random();
            while (true) {
                ValueNumber = rand.Next(0, 20);
                Console.WriteLine("Значение числа =  " + ValueNumber);
                NotifyObservers();
                Thread.Sleep(1000 + rand.Next(0, 500));

            }
        }
    }

}
