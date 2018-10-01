import Data.List

-- Wrapper type to describe each single task
newtype Task = Task { task :: String } deriving(Eq)

-- Each two lines will be 
data TaskPair = TaskPair { t1 :: Task
                         , t2 :: Task 
                         }

convertToTaskPair :: String -> String -> TaskPair
convertToTaskPair s1 s2 = 
    let t1 = Task s1
        t2 = Task s2
    in TaskPair t1 t2

getTaskPairList :: [String] -> [TaskPair]
getTaskPairList [] = []
getTaskPairList (s1:s2:s) = convertToTaskPair s1 s2 : getTaskPairList s

getFirstTask :: TaskPair -> Task
getFirstTask = t1 

getSecondTask :: TaskPair -> Task
getSecondTask = t2

taskToString :: Task -> String
taskToString = task

taskPairToString :: TaskPair -> String
taskPairToString tp = 
    let t1 = getFirstTask tp
        t2 = getSecondTask tp
        s1 = "Task 1: " ++ taskToString t1 
        s2 = "Task 2: " ++ taskToString t2
    in (s1 ++ s2)

main = do 
    get_input <- getContents
    let all_input = lines get_input
    let taskList = getTaskPairList all_input
    mapM taskList
