curl -s http://localhost:8080/dotc/dotc/UpdateBlockInfo -X POST -H "Content-Type:application/json" -d '{"limit":"5"}' >> dotc-update.out 2>&1
echo " -"  $(date "+%Y-%m-%d %H:%M:%S") >> dotc-update.out