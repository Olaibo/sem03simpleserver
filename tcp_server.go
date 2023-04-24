package main

import(
	"io"
	"log"
	"net"
	"sync"
	"github.com/Olaibo/is105sem03/mycrypt"
	"github.com/Trym123/minyr/yr"
)

func main() {

	varwgsync.WaitGroup

	server, err:=net.Listen("tcp", "172.17.0.3:38527")
	iferr!=nil{
		log.Fatal(err)
	}
	log.Printf("bundet til %s", server.Addr().String())
	wg.Add(1)
	gofunc() {
		deferwg.Done()
		for{
			log.Println("før server.Accept() kallet")
			conn, err:=server.Accept()
			iferr!=nil{
				return
			}
			gofunc(cnet.Conn) {
				deferc.Close()
				for{
					buf:=make([]byte, 1024)
					n, err:=c.Read(buf)
					iferr!=nil{
						iferr!=io.EOF{
							log.Println(err)
						}
						return// fra for løkke
					}

					dekryptertMelding:=mycrypt.Krypter([]rune(string(buf[:n])), mycrypt.ALF_SEM03, len(mycrypt.ALF_SEM03)-4)
					log.Println("Dekrypter melding: ", string(dekryptertMelding))
					switchmsg:=string(dekryptertMelding); msg{
  				        case"ping":
						//_, err = c.Write([]byte("pong"))
						kryptertMelding:=mycrypt.Krypter([]rune(string("pong")), mycrypt.ALF_SEM03, 4)
						log.Println("krypter melding: ", string(kryptertMelding))
						_, err=conn.Write([]byte(string(kryptertMelding)))
					case"Kjevik;SN39040;18.03.2022 01:50;6":
						newString, err:=yr.CelsiusToFahrenheitLine("Kjevik;SN39040;18.03.2022 01:50;6")
							iferr!=nil{
							log.Fatal(err)
							}
							kryptertMelding:=mycrypt.Krypter([]rune(string(newString)), mycrypt.ALF_SEM03, 4)
							_, err=conn.Write([]byte(string(kryptertMelding)))

					default:
						_, err=c.Write(buf[:n])
					}
					iferr!=nil{
						iferr!=io.EOF{
							log.Println(err)
						}
						return// fra for løkke
					}
				}
			}(conn)
		}
	}()
	wg.Wait()
}
