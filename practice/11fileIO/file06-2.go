package main
import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
)
/*
	subject : ls付き　サンプルで学ぶ Go 言語：Temporary Files and Directories
*/
func check(e error) {
	if e != nil {
		panic(e)
	}
}
func ls(dir string) {
	out, err := exec.Command("ls", "-laF", dir).Output()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(out))
	}
}

func main() {

	//The easiest way to create a temporary file is by calling ioutil.TempFile.
	//It creates a file and opens it for reading and writing. We provide "" as the first argument,
	//so ioutil.TempFile will create the file in the default location for our OS.
	f, err := ioutil.TempFile("", "sample")
	check(err)

	//Display the name of the temporary file. On Unix-based OSes the directory will likely be /tmp.
	//The file name starts with the prefix given as the second argument to ioutil.TempFile
	//and the rest is chosen automatically to ensure that concurrent calls will always create different file names.
	fmt.Println("Temp file name:", f.Name())
	ls(f.Name())
	
	//Clean up the file after we’re done.
	//The OS is likely to clean up temporary files by itself after some time, but it’s good practice to do this explicitly.
	defer os.Remove(f.Name())

	//We can write some data to the file.
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)
	ls(f.Name())

	//If we intend to write many temporary files, we may prefer to create a temporary directory.
	//ioutil.TempDir’s arguments are the same as TempFile’s, but it returns a directory name rather than an open file.
	dname, err := ioutil.TempDir("", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname)
	ls(dname + "/..")
	
	defer os.RemoveAll(dname)

	//Now we can synthesize temporary file names by prefixing them with our temporary directory.
	fname := filepath.Join(dname, "file1")
	err = ioutil.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
	ls(dname)
}
// -*- mode: compilation; default-directory: "~/go/src/practice/11fileIO/" -*-
// Compilation started at Tue Oct 19 21:03:21
//  
// go run file06-2.go
// Temp file name: /var/folders/5m/29zwdxmj52q7klt6qnkzxn_40000gp/T/sample992448840
// -rw-------  1 Akira  staff  0 10 19 21:03 /var/folders/5m/29zwdxmj52q7klt6qnkzxn_40000gp/T/sample992448840
//  
// -rw-------  1 Akira  staff  4 10 19 21:03 /var/folders/5m/29zwdxmj52q7klt6qnkzxn_40000gp/T/sample992448840
//  
// Temp dir name: /var/folders/5m/29zwdxmj52q7klt6qnkzxn_40000gp/T/sampledir828113927
// total 16
// drwx------    2 Akira  staff    64  9 22 23:52 ${DaemonNameOrIdentifierHere}/
// drwx------@ 138 Akira  staff  4416 10 19 21:03 ./
// drwxr-xr-x@   7 Akira  staff   224  9 22 23:45 ../
// drwx------    5 Akira  staff   160 10 16 03:48 .AddressBookLocks/
// drwx------    2 Akira  staff    64  9 22 23:52 .CalendarLocks/
// drwxr-xr-x    9 Akira  staff   288 10 19 20:48 .LINKS/
// -rw-r--r--    1 Akira  staff     0 10 16 03:38 .keystoneAgentLock
// drwxr-xr-x    2 Akira  staff    64  9 26 03:39 9890E4EA-D6C3-4D0C-BC51-E0374EA11E03/
// drwx------    2 Akira  staff    64  9 22 23:54 AudioComponentRegistrar/
// drwxr-xr-x    2 Akira  staff    64  9 26 03:39 EBD1692F-EA92-4DCA-B915-9DFDF88311D9/
// drwx------@   2 Akira  staff    64 10 13 21:57 Temp-aa73942f-d25c-2043-a3ee-93b7a38f8e8c/
// drwx------@   2 Akira  staff    64 10 19 21:01 TemporaryItems/
// drwx------@   2 Akira  staff    64  9 22 23:52 com.adobe.accmac.ACCFinderSync/
// drwx------@   2 Akira  staff    64  9 23 13:25 com.apple.AMPArtworkAgent/
// drwx------@   2 Akira  staff    64  9 22 23:52 com.apple.AMPDeviceDiscoveryAgent/
// drwx------@   2 Akira  staff    64  9 22 23:52 com.apple.AddressBook.ContactsAccountsService/
// drwx------@   2 Akira  staff    64  9 22 23:52 com.apple.AirPlayUIAgent/
// drwx------    2 Akira  staff    64  9 22 23:52 com.apple.AppSSOAgent/
// drwx------@   3 Akira  staff    96  9 23 13:25 com.apple.BKAgentService/
// drwx------@   3 Akira  staff    96  9 22 23:52 com.apple.CalendarAgent/
// drwx------@   3 Akira  staff    96  9 22 23:52 com.apple.CalendarNotification.CalNCService/
// drwx------    2 Akira  staff    64  9 23 00:51 com.apple.CloudDocs.MobileDocumentsFileProvider/
// drwx------    3 Akira  staff    96 10  8 03:25 com.apple.CloudDocsDaemon.container-metadata-extractor/
// drwx------@   3 Akira  staff    96  9 22 23:52 com.apple.CloudPhotosConfiguration/
// drwx------@   2 Akira  staff    64 10 19 16:48 com.apple.ContactsAgent/
// drwx------@   2 Akira  staff    64  9 23 01:52 com.apple.CoreRoutine.helperservice/
// drwx------@   2 Akira  staff    64  9 22 23:52 com.apple.CryptoTokenKit.pivtoken/
// drwx------@   2 Akira  staff    64  9 22 23:52 com.apple.CryptoTokenKit.setoken/
// drwx------@   3 Akira  staff    96  9 26 12:07 com.apple.DataDetectorsLocalSources/
// drwx------    2 Akira  staff    64  9 22 23:54 com.apple.ImageIOXPCService/
// drwx------@   2 Akira  staff    64  9 22 23:52 com.apple.LoginUserService/
// drwx------@   2 Akira  staff    64  9 23 11:18 com.apple.MailCacheDelete/
// drwx------@   4 Akira  staff   128 10  7 16:35 com.apple.Maps/
// drwx------@   2 Akira  staff    64  9 22 23:52 com.apple.Music.MusicCacheExtension/
// drwx------@   3 Akira  staff    96  9 24 15:22 com.apple.Notes/
// drwx------@   2 Akira  staff    64  9 22 23:52 com.apple.Notes.IntentsExtension/
// drwx------@   2 Akira  staff    64  9 22 23:52 com.apple.Notes.WidgetExtension/
// drwx------@   2 Akira  staff    64  9 23 00:46 com.apple.OSDUIHelper/
// drwx------@   4 Akira  staff   128 10  1 13:53 com.apple.Photos/
// drwx------@   2 Akira  staff    64  9 22 23:52 com.apple.Photos.PhotosReliveWidget/
// drwx------@   4 Akira  staff   128 10  7 20:28 com.apple.Preview/
// drwx------@   5 Akira  staff   160  9 22 23:55 com.apple.Safari/
// drwx------@   2 Akira  staff    64  9 22 23:54 com.apple.Safari.BrowserDataImportingService/
// drwx------@   3 Akira  staff    96  9 22 23:52 com.apple.Safari.CacheDeleteExtension/
// drwx------    2 Akira  staff    64  9 22 23:54 com.apple.SafariLaunchAgent/
// drwx------@   3 Akira  staff    96  9 22 23:52 com.apple.ScreenTimeAgent/
// drwx------@   2 Akira  staff    64  9 22 23:52 com.apple.ScreenTimeWidgetApplication.ScreenTimeWidgetExtension/
// drwx------@   2 Akira  staff    64  9 25 13:56 com.apple.Siri/
// drwx------@   2 Akira  staff    64  9 25 13:56 com.apple.SiriUI.SiriUISetupXPC/
// drwx------@   2 Akira  staff    64  9 22 23:52 com.apple.TV.TVCacheExtension/
// drwx------    3 Akira  staff    96  9 22 23:52 com.apple.TelephonyUtilities/
// drwx------@   2 Akira  staff    64  9 22 23:52 com.apple.UsageTrackingAgent/
// drwx------@   2 Akira  staff    64  9 22 23:54 com.apple.accessibility.mediaaccessibilityd/
// drwx------    2 Akira  staff    64  9 22 23:52 com.apple.akd/
// drwx------    2 Akira  staff    64  9 22 23:53 com.apple.amp.mediasharingd/
// drwx------    2 Akira  staff    64  9 22 23:57 com.apple.animoji/
// drwx------    2 Akira  staff    64  9 22 23:53 com.apple.ap.adprivacyd/
// drwx------    3 Akira  staff    96  9 22 23:53 com.apple.appstoreagent/
// drwx------    3 Akira  staff    96  9 22 23:53 com.apple.bird/
// drwx------    3 Akira  staff    96  9 22 23:52 com.apple.cloudd/
// drwxr-xr-x    3 Akira  staff    96  9 22 23:53 com.apple.cloudkit.upload-request.cache/
// drwx------@   3 Akira  staff    96  9 23 01:52 com.apple.contacts.donation-agent/
// drwx------@   2 Akira  staff    64 10  2 21:58 com.apple.corerecents.recentsd/
// drwx------    2 Akira  staff    64  9 22 23:52 com.apple.corespeechd/
// drwx------    2 Akira  staff    64  9 23 13:21 com.apple.devicecheckd/
// drwx------    3 Akira  staff    96  9 22 23:52 com.apple.dmd/
// drwx------@   2 Akira  staff    64  9 23 11:18 com.apple.dt.IDECacheDeleteAppExtension/
// drwx------    2 Akira  staff    64  9 22 23:52 com.apple.fileproviderd/
// drwx------@   3 Akira  staff    96  9 22 23:52 com.apple.geod/
// drwx------@   3 Akira  staff    96  9 22 23:53 com.apple.iCal.CalendarWidgetExtension/
// drwx------    3 Akira  staff    96  9 22 23:52 com.apple.icloud.searchpartyd/
// drwx------    3 Akira  staff    96  9 22 23:52 com.apple.identityservicesd/
// drwx------    2 Akira  staff    64  9 22 23:52 com.apple.imagent/
// drwx------    3 Akira  staff    96  9 22 23:53 com.apple.imdpersistence.IMDPersistenceAgent/
// drwx------@   3 Akira  staff    96  9 22 23:52 com.apple.inputmethod.Kotoeri.KanaTyping/
// drwx------@   2 Akira  staff    64  9 22 23:52 com.apple.languageassetd/
// drwx------@   5 Akira  staff   160 10  2 21:57 com.apple.mail/
// drwx------    2 Akira  staff    64  9 22 23:52 com.apple.mapspushd/
// drwx------@   4 Akira  staff   128 10  8 01:31 com.apple.mediaanalysisd/
// drwx------@   3 Akira  staff    96  9 29 11:15 com.apple.mediastream.mstreamd/
// drwx------@   2 Akira  staff    64  9 23 00:00 com.apple.mobiletimer.WorldClockWidget/
// drwx------@   3 Akira  staff    96  9 22 23:52 com.apple.notificationcenterui/
// drwx------    3 Akira  staff    96  9 22 23:52 com.apple.nsurlsessiond/
// drwx------    3 Akira  staff    96  9 23 00:22 com.apple.parsec-fbf/
// drwx------@   3 Akira  staff    96 10 19 18:17 com.apple.parsecd/
// drwx------    3 Akira  staff    96  9 22 23:52 com.apple.passd/
// drwx------@   4 Akira  staff   128 10 19 02:46 com.apple.photoanalysisd/
// drwx------@   4 Akira  staff   128 10 18 11:03 com.apple.photolibraryd/
// drwx------@   2 Akira  staff    64 10  8 01:30 com.apple.photos.ImageConversionService/
// drwx------    3 Akira  staff    96  9 22 23:52 com.apple.pluginkit/
// drwx------@   2 Akira  staff    64  9 22 23:57 com.apple.preferencepane.security.AdvertisingExtension/
// drwx------@   2 Akira  staff    64  9 22 23:57 com.apple.preferencepane.security.PrivacyAnalytics/
// drwx------@   2 Akira  staff    64  9 22 23:57 com.apple.preferencepane.security.PrivacyTrackingAwareness.PrivacyPhotos/
// drwx------    2 Akira  staff    64  9 28 17:09 com.apple.printtool.agent/
// drwx------    2 Akira  staff    64  9 23 13:21 com.apple.proactiveeventtrackerd/
// drwx------    2 Akira  staff    64  9 22 23:52 com.apple.progressd/
// drwx------@   2 Akira  staff    64  9 22 23:52 com.apple.quicklook.QuickLookUIService/
// drwx------    3 Akira  staff    96  9 22 23:52 com.apple.quicklook.ThumbnailsAgent/
// drwx------    2 Akira  staff    64  9 22 23:52 com.apple.quicklook.satellite.general/
// drwx------@   2 Akira  staff    64  9 23 13:56 com.apple.quicklook.ui.helper/
// drwx------    4 Akira  staff   128 10 18 15:03 com.apple.remindd/
// drwx------    2 Akira  staff    64  9 22 23:52 com.apple.replayd/
// drwx------@   3 Akira  staff    96  9 22 23:52 com.apple.routined/
// drwx------    4 Akira  staff   128  9 22 23:52 com.apple.sharingd/
// drwx------    2 Akira  staff    64  9 25 13:56 com.apple.siri-distributed-evaluation/
// drwx------@   2 Akira  staff    64  9 23 13:25 com.apple.siri.media-indexer/
// drwx------    2 Akira  staff    64  9 23 00:47 com.apple.speech.speechsynthesisd.x86_64/
// drwx------@   4 Akira  staff   128  9 22 23:52 com.apple.stocks.detailintents/
// drwx------@   4 Akira  staff   128  9 22 23:52 com.apple.stocks.widget/
// drwx------    4 Akira  staff   128  9 22 23:53 com.apple.studentd/
// drwx------    2 Akira  staff    64  9 22 23:52 com.apple.tccd/
// drwx------    3 Akira  staff    96  9 22 23:52 com.apple.touristd/
// drwx------    2 Akira  staff    64  9 23 04:52 com.apple.transparencyd/
// drwx------    3 Akira  staff    96  9 22 23:52 com.apple.triald/
// drwx------    3 Akira  staff    96  9 22 23:44 com.apple.trustd/
// drwx------    3 Akira  staff    96  9 23 00:47 com.apple.useractivityd/
// drwx------@   2 Akira  staff    64  9 22 23:52 com.apple.weather.WeatherIntents/
// drwx------@   3 Akira  staff    96  9 22 23:52 com.apple.weather.widget/
// drwxr-xr-x    2 Akira  staff    64  9 22 23:52 com.apple.wifivelocity/
// drwx------@   2 Akira  staff    64  9 22 23:52 com.kensington.tbwdkmanager/
// drwx------@   4 Akira  staff   128 10 18 22:50 com.tencent.xinWeChat/
// drwx------@   2 Akira  staff    64 10 18 22:50 com.tencent.xinWeChat.MiniProgram/
// drwx------    2 Akira  staff    64  9 22 23:52 diagnosticextensionsd/
// drwx------    3 Akira  staff    96 10 15 19:07 go-build193554147/
// drwx------    3 Akira  staff    96 10  6 22:52 go-build3536493250/
// drwx------    3 Akira  staff    96 10 19 21:03 go-build4060088483/
// drwx------    3 Akira  staff    96  9 29 15:50 go-build677834319/
// drwx------    3 Akira  staff    96  9 22 23:52 homed/
// drwx------    3 Akira  staff    96  9 22 23:52 itunescloudd/
// drwxr-xr-x    2 Akira  staff    64 10 19 16:48 journeys/
// -rw-------    1 Akira  staff     4 10 19 21:03 sample992448840
// drwx------    2 Akira  staff    64 10 19 21:03 sampledir828113927/
// drwxr-xr-x    2 Akira  staff    64  9 22 23:53 skyeTemp/
// drwx------    2 Akira  staff    64  9 22 23:52 studentd/
// srwxr-xr-x    1 Akira  staff     0  9 23 00:49 vscode-git-47140e43d8.sock=
// srwxr-xr-x    1 Akira  staff     0  9 23 00:56 vscode-git-5ed9597399.sock=
// srwxr-xr-x    1 Akira  staff     0 10 18 14:03 vscode-git-c47b4175d3.sock=
// -rw-------    1 Akira  staff  2169 10  9 17:17 xcrun_db
//  
// total 8
// drwx------    3 Akira  staff    96 10 19 21:03 ./
// drwx------@ 138 Akira  staff  4416 10 19 21:03 ../
// -rw-r--r--    1 Akira  staff     2 10 19 21:03 file1
//  
//  
// Compilation finished at Tue Oct 19 21:03:21
