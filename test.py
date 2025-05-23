import python.package as bcd
import bedrockworldoperator as bwo

c = bwo.new_chunk()
for i in c.sub():
    i.set_block(0, 0, 0, 0, bwo.AIR_BLOCK_RUNTIME_ID)

sub_paylods = []
for index, value in enumerate(c.sub()):
    temp = bwo.sub_chunk_disk_payload(value, index)
    sub_paylods.append(temp)

db = bcd.new_timeline_database("ss")
print(db.is_valid())

tl = db.new_chunk_timeline(bcd.ChunkPos(0, 1))
print(tl.is_valid())
print(tl.all_time_point())

tl.append_disk_chunk(bcd.ChunkData(sub_paylods, []))
print(tl.all_time_point())
tl.save()

db.close_timeline_db()
