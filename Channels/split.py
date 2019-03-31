import os
path1= "C:/just_trial/ObjectCategories/training_set"
path2 = "C:/just_trial/ObjectCategories/test_set"

directory = "C:/just_trial/ObjectCategories"
fileName = os.listdir(directory)
for subdir , dir , files in os.walk(directory) :
    print (subdir,"\t",files)
    split1 = int(0.8*len(files))
    split2=int(0.8*len(files))
    train_files = files[:split1]
    test_files = files[split2:]
    for iter in range(len(fileName)) :
        for file in train_files :
            print(os.path.abspath(file))


