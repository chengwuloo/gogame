```
��ʵ�ַ�����ʽ��������Ϸ�첽��ܣ���˼ܹ���

��  main.go                                ���������
��  README.md                              �ĵ�˵��
��  
����comm                                    ͨ�ú���
��  ����utils
��          utils.go
��          
����core                                    �������
��  ��  calback.go                          �ص�����
��  ��  mailbox.go                          �ʲ۹���
��  ��  msq.go                              ��Ϣ����
��  ��  proc.go                             ��Ϣ������(EventLoop)
��  ��  slot.go                             �ʲ�(EventLoopThread)
��  ��  worker.go                           ҵ����
��  ��  
��  ����timerv2                              ��ʱ��
��          cron_test.go
��          hashwheel.go
��          README.md
��          safetimer.go
��          safetimer_test.go
��          timer.go
��          
����public                                  ��������
��  ����define
��  ��      define.go                       ��������
��  ��      
��  ����iface                                �ӿ�
��      ��  idesk.go                        ���ӽӿ�
��      ��  idesk_delegate.go               ��Ϸ���ӽӿ�      --����Ϸ�����߼��̳�ʹ��
��      ��  iplayer.go                      ��ҽӿ�
��      ��  irobot_delegate.go              ����Ϸ�����˽ӿ�  --����Ϸ�����˼̳�ʹ��
��      ��  
��      ����impl                             ʵ��
��              desk.go                    ������
��              desk_mgr.go                ���ӹ�����
��              player.go                  �����
��              player_mgr.go              ��ҹ�����
��              robot.go                   ��������
��              robot_mgr.go               �����˹�����
��              
����service                                 �������
��      sentry.go                          ��Ӧÿ������ҵ����Ϣ���
��      smain.go                           ��Ӧÿ������ҵ���߼����� - ����Ϸ���ʵ��
��      
����sub                                     ��������Ϸģ��           - ����Ϸ����ʵ��
    ����game_dragon_tiger                   ������
            desk.go                       ����Ϸ������
            robot.go                      ����Ϸ��������
            
